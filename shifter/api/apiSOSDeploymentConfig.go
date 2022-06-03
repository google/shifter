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
