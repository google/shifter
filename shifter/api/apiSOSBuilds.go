/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/openshift/api/build/v1"
	"log"
	"net/http"
	os "shifter/openshift"
)

type SOSBuild struct {
	Shifter Shifter  `json:"shifter"`
	Build   v1.Build `json:"build"`
}

type SOSBuilds struct {
	Shifter Shifter      `json:"shifter"`
	Builds  v1.BuildList `json:"builds"`
}

func (server *Server) SOSGetBuildsByProject(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSBuilds SOSBuilds
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSBuilds)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSBuilds.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSBuilds.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	builds, err := openshift.GetAllBuilds(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// Add Projects to the Response
	sOSBuilds.Builds = *builds
	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSBuilds)
}

func (server *Server) SOSGetBuild(ctx *gin.Context) {

	// Validate Project Name has been Provided
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Validate Route Name has been Provided
	buildName := ctx.Param("buildName")
	if buildName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Build Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSBuild SOSBuild
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSBuild)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSBuild.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSBuild.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSBuild.Shifter.ClusterConfig.Username
	openshift.Password = sOSBuild.Shifter.ClusterConfig.Password

	build, err := openshift.GetBuild(projectName, buildName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Route to the Response
	sOSBuild.Build = *build

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSBuild)
}
