/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	generator "shifter/generator"
	lib "shifter/lib"
	os "shifter/openshift"
	ops "shifter/ops"
	"shifter/processor"

	"github.com/gin-gonic/gin"
	osNativeDC "github.com/openshift/api/apps/v1"
	osNativeProject "github.com/openshift/api/project/v1"
)

type Shifter struct {
	ClusterConfig *ClusterConfig `json:"clusterConfig"`
}

type Convert struct {
	Shifter *Shifter       `json:"shifter"`
	Items   []*ConvertItem `json:"items"`
}

type ResponseConvert struct {
	SUID    ops.SUID `json:"suid"`
	Message string   `json:"message"`
}

type ConvertItem struct {
	Namespace        *osNativeProject.Project     `json:"namespace"`
	DeploymentConfig *osNativeDC.DeploymentConfig `json:"deploymentConfig"`
	// Options * ConvertOptions `json:"options"`
}

type ClusterConfig struct {
	ConnectionName string `json:"connectionName"`
	BaseUrl        string `json:"baseUrl"`
	BearerToken    string `json:"bearerToken"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}

func (server *Server) Convert(ctx *gin.Context) {

	// Create API Unique RUN ID
	suid := ops.CreateSUID("") //NESTED TODO - Error Handling

	// Instanciate a Shifter Convert Structure
	convert := Convert{}
	// using BindJSON method to serialize API Request Body with struct
	if err := ctx.BindJSON(&convert); err != nil {
		// Error: Unable to Parse Request JSON -> Convert Struct
		log.Printf("🌐 ❌ ERROR: Unable to Parse API Request JSON -> Convert Struct, Returning: Status %d.", http.StatusBadRequest)
		// Return Error JSON Response to API Call
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO - Turn this into a Global Function called at the beginign of all API Calls with Validations and Responses
	var openshift os.Openshift
	openshift.Endpoint = convert.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = convert.Shifter.ClusterConfig.BearerToken
	openshift.Username = convert.Shifter.ClusterConfig.Username
	openshift.Password = convert.Shifter.ClusterConfig.Password

	// Check that API Request contains Items for Conversion
	if len(convert.Items) <= 0 {
		// Error: No items provided for conversion.
		log.Printf("🌐 ❌ ERROR: No items have been provided for conversion. Returning: Status %d.", http.StatusBadRequest)
		// Create Error
		err := errors.New("No items have been provided for conversion")
		// Return Error JSON Response to API Call
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else {
		/*
			Process Conversion Items
			Details: When this API endpoint is called by a client, it will contain an array of items
			that need to be converted and returned.
		*/
		for _, item := range convert.Items {

			// Validate that the provided namespace is valid in the cluster.
			deploymentConfig, err := openshift.GetDeploymentConfig(item.Namespace.ObjectMeta.Name, item.DeploymentConfig.ObjectMeta.Name)
			// TODO - BULK Error Catch [Long Comment EOF]
			if err != nil {
				// Error validating provided OpenShift Namespace
				log.Printf("🌐 ❌ ERROR: Unable to locate the provided requested OpenShift DeploymentConfig: %s within OpenShift Namespace: %s", item.DeploymentConfig.ObjectMeta.Name, item.Namespace.ObjectMeta.Name)
				// Return Error JSON Response to API Call
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}

			// JSON Marshal the Contents of the OpenShift DeploymentConfig Object to JSON
			osDeploymentConfig, err := json.Marshal(deploymentConfig)
			if err != nil {
				// Error: Unable to Parse OpenShift DeploymentConfig Object to JSON object
				log.Printf("🌐 ❌ ERROR: Unable to Parse OpenShift DeploymentConfig Object -> JSON, Returning: Status %d.", http.StatusBadRequest)
				// Return Error JSON Response to API Call
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}

			// Handle the Conversion of the Manifests and File Writing
			//var generator generator.Generator
			var objs []lib.K8sobject
			//NESTED TODO - Error Handling in Processor Call
			obj := processor.Processor(osDeploymentConfig, "DeploymentConfig", nil)
			for _, v := range obj {
				objs = append(objs, v)
			}

			// Log Info
			log.Printf("🌐 💡 INFO: Preparing to Convert %d OpenShift DeploymentConfigs", len(objs))

			//NESTED TODO - Error Handling in New Generator Call
			convertedObjects := generator.NewGenerator("yaml", item.DeploymentConfig.ObjectMeta.Name, objs)
			// Loop Through Each ConvertedObject and create File Object
			for _, conObj := range convertedObjects {
				// Create File Object
				fileObj := &ops.FileObject{
					StorageType:   server.config.serverStorage.storageType,
					Path:          (server.config.serverStorage.sourcePath + "/" + suid.DirectoryName + "/" + item.Namespace.ObjectMeta.Name + "/" + item.DeploymentConfig.ObjectMeta.Name),
					Ext:           "yaml",
					Content:       conObj.Payload,
					ContentLength: conObj.Payload.Len(),
				}
				// Log FileObject
				fileObj.Meta()
				//NESTED TODO - Error Handling in Write File Call
				fileObj.WriteFile()
			}
		}

		// Zip / Package Converted Objects
		//NESTED TODO - Error Handling in Archive Call
		err := ops.Archive(server.config.serverStorage.sourcePath, server.config.serverStorage.outputPath, suid)
		if err != nil {
			// Error: Unable to Archive Directory of Objects
			log.Printf("🌐 ❌ ERROR: Unable to Archive Directory of Objects, Returning: Status %d.", http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}

		// API Convert Successful
		log.Printf("✅ SUCCESS: API Convert - %d Objects Converted", len(convert.Items))
		// Return API JSON Response
		ctx.JSON(
			http.StatusOK,
			// Construct API Endpoint Response
			ResponseConvert{
				SUID:    suid,
				Message: "Converted " + fmt.Sprint(len(convert.Items)) + " Objects",
			})
	}
}

/*
	----------------------------------------------------
		TODO - BULK Error Catch [Long Comment EOF]
	----------------------------------------------------
	Currently the API call for convert will take a set of items to be converted. The moment one
	of these falls over of fails at an point in "the loop". The API call will terminate in error and return.

	Problems with approach:
		- If you have multiple errors you need to solve each error just to see the next one.
		- Bad user experience.
		- Lack of ability for parallel processing and conversion of objects

	Possible Solution:
		- Construct an error array and allow the loop to complete converting as many objects as possible
		- Catching all the conversion/lookup/validation records along the way only converting successful objects
		- Returning successful objects as planned in a downloadable file.
		- Also enabling us to write out conversion logs and provide them as part of the archive.

	Best Long term Solution:
		- All of the above +
		- Factory in Object by Object customization (Object level Flags, Route Changes, Container Re-Writes/tags)
		- Convert to background job with the ability to list running jobs on the UI with link to completed download archive.
*/
