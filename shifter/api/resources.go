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
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/types"
	"net/http"
	"shifter/lib"
	openshift "shifter/openshift"
)

type Resource struct {
	Shifter   Shifter `json:"shifter"`
	Resources *[]openshift.ResourceList
}

func (server *Server) GetResources(ctx *gin.Context) {
	projectName := ctx.Param("projectName")
	resourceName := ctx.Param("resourceName")
	resourceKind := ctx.Param("resourceKind")
	resourceUid := types.UID(ctx.Param("resourceUid"))

	if projectName == "" {
		err := errors.New("Project must be specified")
		lib.CLog("fatal", "Missing project name", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var sOSResources Resource
	err := json.NewDecoder(ctx.Request.Body).Decode(&sOSResources)
	if err != nil {
		lib.CLog("error", "decoding resources", err)
		ctx.JSON(501, gin.H{"error": err})
	}

	var o openshift.Openshift
	o.Endpoint = sOSResources.Shifter.ClusterConfig.BaseUrl
	o.AuthToken = sOSResources.Shifter.ClusterConfig.BearerToken

	resources, err := o.GetResources(projectName, false, resourceKind, resourceName, resourceUid)
	if err != nil {
		lib.CLog("error", "Unable to get resources for "+projectName+" "+resourceKind+"\\"+resourceName, err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	sOSResources.Resources = &resources
	ctx.JSON(http.StatusOK, sOSResources)
}
