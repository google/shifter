/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	os "shifter/openshift"

	"github.com/gin-gonic/gin"
	osNative "github.com/openshift/api/apps/v1"
)

type SOSDeploymentConfig struct {
	Shifter          Shifter                   `json:"shifter"`
	DeploymentConfig osNative.DeploymentConfig `json:"deploymentConfig"`
}

func (server *Server) SOSGetDeploymentConfig(ctx *gin.Context) {

	// Validate Project Name has been Provided
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Validate DeploymentConfig Name has been Provided
	deploymentConfigName := ctx.Param("deploymentConfigName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Deployment Config Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSDeploymentConfig SOSDeploymentConfig
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSDeploymentConfig)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSDeploymentConfig.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSDeploymentConfig.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSDeploymentConfig.Shifter.ClusterConfig.Username
	openshift.Password = sOSDeploymentConfig.Shifter.ClusterConfig.Password

	deploymentconfig, err := openshift.GetDeploymentConfig(projectName, deploymentConfigName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add DeploymentConfig to the Response
	sOSDeploymentConfig.DeploymentConfig = *deploymentconfig

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfig)
}
