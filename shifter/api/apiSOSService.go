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
	v1 "k8s.io/api/core/v1"
	"log"
	"net/http"
	os "shifter/openshift"
)

type SOSService struct {
	Shifter Shifter    `json:"shifter"`
	Service v1.Service `json:"service"`
}

type SOSServices struct {
	Shifter  Shifter        `json:"shifter"`
	Services v1.ServiceList `json:"servicelist"`
}

func (server *Server) SOSGetServicesByProject(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSServices SOSServices
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSServices)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSServices.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSServices.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	services, err := openshift.GetAllServices(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// Add Projects to the Response
	sOSServices.Services = *services
	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSServices)
}

func (server *Server) SOSGetService(ctx *gin.Context) {

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
	serviceName := ctx.Param("serviceName")
	if serviceName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Service Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSService SOSService
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSService)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSService.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSService.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSService.Shifter.ClusterConfig.Username
	openshift.Password = sOSService.Shifter.ClusterConfig.Password

	service, err := openshift.GetService(projectName, serviceName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Route to the Response
	sOSService.Service = *service

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSService)
}
