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

func (server *Server) SOSGetDeploymentConfigsByProject(ctx *gin.Context) {

	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusMisdirectedRequest, errorResponse(err))
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
	openshift := os.NewClient(http.DefaultClient)
	// Configure Authorization
	openshift.AuthOptions = &os.AuthOptions{BearerToken: sOSDeploymentConfigs.Shifter.ClusterConfig.BearerToken}
	openshift.BaseURL, err = url.Parse(sOSDeploymentConfigs.Shifter.ClusterConfig.BaseUrl)
	if err != nil {
		panic(err)
	}

	// Get List of OpenShift Projects
	deploymentconfigs, err := openshift.Apis.DeploymentConfigs.GetByProject(projectName)
	if err != nil {
		fmt.Println(err)
	}

	// Add Projects to the Response
	sOSDeploymentConfigs.DeploymentConfigs = *deploymentconfigs

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSDeploymentConfigs)
}
