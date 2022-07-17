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
	"fmt"
	"net/http"
	os "shifter/openshift"

	"github.com/gin-gonic/gin"
	osNative "github.com/openshift/api/apps/v1"
)

type SOSDeploymentConfigs struct {
	Shifter           Shifter                       `json:"shifter"`
	DeploymentConfigs osNative.DeploymentConfigList `json:"deploymentConfigs"`
}

func (server *Server) SOSGetDeploymentConfigs(ctx *gin.Context) {

	// TODO: Handle Incorrect Payload. Where Shifter{} block or incorrect Shifter{} block is provided in body

	// Parse REST Request JSON Body
	var sOSDeploymentConfigs SOSDeploymentConfigs
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSDeploymentConfigs)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSDeploymentConfigs.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSDeploymentConfigs.Shifter.ClusterConfig.BearerToken

	deploymentconfigs, err := openshift.GetAllDeploymentConfigs("default")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Projects to the Response
	sOSDeploymentConfigs.DeploymentConfigs = *deploymentconfigs

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfigs)
}
