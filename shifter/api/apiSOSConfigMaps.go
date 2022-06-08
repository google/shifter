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

type SOSConfigMap struct {
	Shifter   Shifter      `json:"shifter"`
	ConfigMap v1.ConfigMap `json:"configmap"`
}

type SOSConfigMaps struct {
	Shifter    Shifter          `json:"shifter"`
	ConfigMaps v1.ConfigMapList `json:"configmaplist"`
}

func (server *Server) SOSGetConfigMapsByProject(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	if projectName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift Project Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSConfigMaps SOSConfigMaps
	//decoder :=
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSConfigMaps)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSConfigMaps.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSConfigMaps.Shifter.ClusterConfig.BearerToken

	// Get List of OpenShift Projects
	configmaps, err := openshift.GetAllConfigMaps(projectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	// Add Projects to the Response
	sOSConfigMaps.ConfigMaps = *configmaps
	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSConfigMaps)
}

func (server *Server) SOSGetConfigMap(ctx *gin.Context) {

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
	configmapName := ctx.Param("configmapName")
	if configmapName == "" {
		// UUID param required & not found.
		err := errors.New("OpenShift ConfigMap Name must be supplied")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal(err.Error())
		return
	}

	// Parse REST Request JSON Body
	var sOSConfigMap SOSConfigMap
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSConfigMap)
	if err != nil {
		fmt.Printf("error %s", err)
		ctx.JSON(501, gin.H{"error": err})
		return
	}

	// Create OpenShift Client
	var openshift os.Openshift
	openshift.Endpoint = sOSConfigMap.Shifter.ClusterConfig.BaseUrl
	openshift.AuthToken = sOSConfigMap.Shifter.ClusterConfig.BearerToken
	openshift.Username = sOSConfigMap.Shifter.ClusterConfig.Username
	openshift.Password = sOSConfigMap.Shifter.ClusterConfig.Password

	configmap, err := openshift.GetConfigMap(projectName, configmapName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// Add Route to the Response
	sOSConfigMap.ConfigMap = *configmap

	// Return JSON API Response
	ctx.JSON(http.StatusOK, sOSConfigMap)
}
