package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	os "shifter/openshift/v3_11"

	"github.com/gin-gonic/gin"
)

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
	openshift := os.NewClient(http.DefaultClient)
	// Configure Authorization
	openshift.AuthOptions = &os.AuthOptions{BearerToken: sOSDeploymentConfig.Shifter.ClusterConfig.BearerToken}
	openshift.BaseURL, err = url.Parse(sOSDeploymentConfig.Shifter.ClusterConfig.BaseUrl)
	if err != nil {
		panic(err)
	}

	// Get List of OpenShift Projects
	deploymentconfig, err := openshift.Apis.DeploymentConfig.Get(projectName, deploymentConfigName)
	if err != nil {
		fmt.Println(err)
	}

	// Add Project to the Response
	sOSDeploymentConfig.DeploymentConfig = *deploymentconfig

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfig)
}
