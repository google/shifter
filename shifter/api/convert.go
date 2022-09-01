// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	//"encoding/json"
	"errors"
	"net/http"
	"shifter/generator"
	"shifter/lib"
	"shifter/openshift"
	"shifter/ops"
	"shifter/processor"

	"github.com/gin-gonic/gin"
	projectv1 "github.com/openshift/api/project/v1"
)

type Convert struct {
	Shifter *Shifter       `json:"shifter"`
	Items   []*ConvertItem `json:"items"`
}

type Shifter struct {
	ClusterConfig *ClusterConfig `json:"clusterConfig"`
}

type ClusterConfig struct {
	ConnectionName string `json:"connectionName"`
	BaseUrl        string `json:"baseUrl"`
	BearerToken    string `json:"bearerToken"`
}

type ConvertItem struct {
	Namespace *projectv1.Project     `json:"namespace"`
	Resource  openshift.ResourceList `json:"resource"`
	// Options * ConvertOptions `json:"options"`
}

type ResponseConvert struct {
	SUUID   ops.SUUID `json:"suid"`
	Message string    `json:"message"`
}

type ConvertedFile struct {
	Namespace string
	Resource  []lib.Converted
}

func (server *Server) Convert(ctx *gin.Context) {

	// Create API Unique RUN ID
	suuid := ops.CreateUUID("") //NESTED TODO - Error Handling

	// Log SUID
	suuid.Meta()

	// Instanciate a Shifter Convert Structure
	convert := Convert{}
	if err := ctx.BindJSON(&convert); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		lib.CLog("error", "Unable to parse api request json to convert resources -> returning http response "+string(http.StatusBadRequest), err)
		return
	}

	var o openshift.Openshift
	o.Endpoint = convert.Shifter.ClusterConfig.BaseUrl
	o.AuthToken = convert.Shifter.ClusterConfig.BearerToken

	var (
		ns        = map[string]bool{}
		resources []openshift.ResourceList
		converted []ConvertedFile
	)

	if len(convert.Items) == 0 {
		err := errors.New("No resources selected for conversion")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		lib.CLog("error", "Convert api", err)
		return
	}

	for _, item := range convert.Items {
		ns[item.Resource.Namespace] = true
	}

	for namespace, _ := range ns {
		r, _ := o.GetResources(namespace, false, "", "", "")
		for _, v := range r {
			resources = append(resources, v)
		}
	}

	for _, item := range convert.Items {
		for _, r := range resources {
			if item.Resource.UID == r.UID {
				lib.CLog("info", "Converting object "+item.Resource.Namespace+"\\"+item.Resource.Name+" of kind "+item.Resource.Kind)
				c, err := processor.Processor(r.Payload.Bytes(), r.Kind, nil)
				if err != nil {
					lib.CLog("error", "Converting object "+item.Resource.Namespace+"\\"+item.Resource.Name+" of kind "+item.Resource.Kind, err)
					ctx.JSON(http.StatusBadRequest, errorResponse(err))
					break
				}
				cr, err := generator.NewGenerator("yaml", item.Resource.Name, c)
				if err != nil {
					lib.CLog("error", "Converting object "+item.Resource.Namespace+"\\"+item.Resource.Name+" of kind "+item.Resource.Kind, err)
					ctx.JSON(http.StatusBadRequest, errorResponse(err))
					break
				}
				var conv ConvertedFile
				conv.Namespace = item.Resource.Namespace
				conv.Resource = cr
				converted = append(converted, conv)
			}
		}
	}

	for _, c := range converted {
		for _, r := range c.Resource {
			file := &ops.FileObject{
				StorageType:   server.config.serverStorage.storageType,
				Path:          (server.config.serverStorage.outputPath + "/" + c.Namespace + "/" + r.Name),
				Ext:           "yaml",
				Content:       r.Payload,
				ContentLength: r.Payload.Len(),
			}
			err := file.WriteFile()
			if err != nil {
				lib.CLog("error", "Writing file", err)
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
			}
		}
	}
	return
}

/*



		// Zip / Package Converted Objects
		//NESTED TODO - Error Handling in Archive Call
		err := ops.Archive(server.config.serverStorage.sourcePath, server.config.serverStorage.outputPath, suid)
		if err != nil {
			// Error: Unable to Archive Directory of Objects
			log.Printf("🌐 ❌ ERROR: Unable to Archive Directory of Objects, Returning: Status %d.", http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		} else {
			// Succes: Archived Directory of Objects
			log.Printf("🌐 ✅ SUCCESS: Archived Directory of Converted Objects.")
		}

		// API Convert Endpoint Successful
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
*/
// API Convert Endpoint Completed

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
