package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	os "shifter/openshift/v3_11"

	"github.com/gin-gonic/gin"
)

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
	openshift := os.NewClient(http.DefaultClient)
	// Configure Authorization
	fmt.Println(sOSDeploymentConfigs)
	openshift.AuthOptions = &os.AuthOptions{BearerToken: sOSDeploymentConfigs.Shifter.ClusterConfig.BearerToken}
	openshift.BaseURL, err = url.Parse(sOSDeploymentConfigs.Shifter.ClusterConfig.BaseUrl)
	if err != nil {
		panic(err)
	}

	// Get List of OpenShift Projects
	deploymentconfigs, err := openshift.Apis.DeploymentConfigs.Get()
	if err != nil {
		fmt.Println(err)
	}

	// Add Projects to the Response
	sOSDeploymentConfigs.DeploymentConfigs = *deploymentconfigs

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfigs)
}
