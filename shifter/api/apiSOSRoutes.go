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
	v1 "github.com/openshift/api/route/v1"
	"log"
	"net/http"
	os "shifter/openshift"
)

type SOSRoute struct {
	Shifter Shifter  `json:"shifter"`
	Route   v1.Route `json:"route"`
}

type SOSRoutes struct {
	Shifter Shifter      `json:"shifter"`
	Routes  v1.RouteList `json:"routes"`
}

func (server *Server) SOSGetRoutesByProject(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSRoutes SOSRoutes
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSRoutes)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSRoutes.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSRoutes.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	routes, err := openshift.GetAllRoutes(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// Add Projects to the Response
	sOSRoutes.Routes = *routes
	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSRoutes)
}

func (server *Server) SOSGetRoute(ctx *gin.Context) {

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
	routeName := ctx.Param("routeName")
	if routeName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Route Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSRoute SOSRoute
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSRoute)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSRoute.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSRoute.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSRoute.Shifter.ClusterConfig.Username
	openshift.Password = sOSRoute.Shifter.ClusterConfig.Password

	route, err := openshift.GetRoute(projectName, routeName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Route to the Response
	sOSRoute.Route = *route

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSRoute)
}
