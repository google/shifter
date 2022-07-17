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
)

func (server *Server) SOSGetDeploymentConfigsByProject(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSDeploymentConfigs SOSDeploymentConfigs
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSDeploymentConfigs)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSDeploymentConfigs.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSDeploymentConfigs.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	deploymentconfigs, err := openshift.GetAllDeploymentConfigs(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// Add Projects to the Response
	sOSDeploymentConfigs.DeploymentConfigs = *deploymentconfigs
	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfigs)
}
