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
	"net/url"
	os "shifter/openshift/v3_11"

	"github.com/gin-gonic/gin"
)

func (server *Server) OpenShiftGetProjects(ctx *gin.Context) {

	// Parse REST Request JSON Body
	var sOSProjects ShifterGetOpenShiftProjects
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSProjects)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	openshift := os.NewClient(http.DefaultClient)
	// Configure Authorization
	openshift.AuthOptions = &os.AuthOptions{BearerToken: sOSProjects.Shifter.ClusterConfig.BearerToken}
	openshift.BaseURL, err = url.Parse(sOSProjects.Shifter.ClusterConfig.BaseUrl)
	if err != nil {
		panic(err)
	}

	// Get List of OpenShift Projects
	projects, err := openshift.Apis.Projects.Get()
	if err != nil {
		fmt.Println(err)
	}

	// Add Projects to the Response
	sOSProjects.Projects = *projects

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSProjects)
}
