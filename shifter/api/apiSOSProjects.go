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
	"fmt"
	"net/http"
	os "shifter/openshift"

	"github.com/gin-gonic/gin"
)

func (server *Server) SOSGetProjects(ctx *gin.Context) {

	// Parse REST Request JSON Body
	var sOSProjects SOSProjects
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSProjects)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSProjects.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSProjects.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSProjects.Shifter.ClusterConfig.Username
	openshift.Password = sOSProjects.Shifter.ClusterConfig.Password

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
