// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	os "shifter/openshift"

	"github.com/gin-gonic/gin"
	osNativeProject "github.com/openshift/api/project/v1"
)

type Projects struct {
	Shifter  Shifter                     `json:"shifter"`
	Projects osNativeProject.ProjectList `json:"projects"`
}

type Project struct {
	Shifter Shifter                 `json:"shifter"`
	Project osNativeProject.Project `json:"project"`
}

// API Endpoints to get the OpenShift project list and project detail
func (server *Server) GetProjects(ctx *gin.Context) {

	// Parse REST Request JSON Body
	var sOSProjects Projects
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSProjects)
	if err != nil {
		log.Println("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSProjects.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSProjects.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	projects, err := openshift.GetAllProjects()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Projects to the Response
	sOSProjects.Projects = *projects

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSProjects)
}

func (server *Server) GetProject(ctx *gin.Context) {

	// Validate Project Name has been Provided
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Parse REST Request JSON Body
	var sOSProject Project
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSProject)
	if err != nil {
		log.Println("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSProject.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSProject.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	project, err := openshift.GetProject(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Add Project to the Response
	sOSProject.Project = *project

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSProject)
}
