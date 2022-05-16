package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	os "shifter/openshift"
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
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Validate DeploymentConfig Name has been Provided
	deploymentConfigName := ctx.Param("deploymentConfigName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Deployment Config Name must be supplied")
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
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
	deploymentconfig := openshift.GetDeploymentConfig(projectName, deploymentConfigName)

	// Add DeploymentConfig to the Response
	sOSDeploymentConfig.DeploymentConfig = *deploymentconfig

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfig)
}
