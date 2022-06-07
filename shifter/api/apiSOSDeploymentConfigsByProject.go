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
	openshift.Username = sOSDeploymentConfigs.Shifter.ClusterConfig.Username
	openshift.Password = sOSDeploymentConfigs.Shifter.ClusterConfig.Password

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
