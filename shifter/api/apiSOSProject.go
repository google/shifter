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
	"log"
	"net/http"
	os "shifter/openshift"

	"github.com/gin-gonic/gin"
	osNativeProject "github.com/openshift/api/project/v1"
)

type SOSProject struct {
	Shifter Shifter                 `json:"shifter"`
	Project osNativeProject.Project `json:"project"`
}

func (server *Server) SOSGetProject(ctx *gin.Context) {

	// Validate Project Name has been Provided
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSProject SOSProject
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSProject)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSProject.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSProject.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSProject.Shifter.ClusterConfig.Username
	openshift.Password = sOSProject.Shifter.ClusterConfig.Password

	// Get List of OpenShift Projects
	project, err := openshift.GetProject(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Project to the Response
	sOSProject.Project = *project

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSProject)
}
