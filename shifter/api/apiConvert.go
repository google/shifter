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
	"reflect"
	os "shifter/openshift/v3_11"

	"shifter/processor"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) Convert(ctx *gin.Context) {

	// Create API Unique RUN ID
	uuid := uuid.New().String()

	//body:=Body{}
	convert := Convert{}
	// using BindJson method to serialize body with struct
	if err := ctx.BindJSON(&convert); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Process Each Item
	var count int = 0
	for idx, item := range convert.Items {
		// Create OpenShift Client
		openshift := os.NewClient(http.DefaultClient)
		// Configure Authorization
		openshift.AuthOptions = &os.AuthOptions{
			BearerToken: convert.Shifter.ClusterConfig.BearerToken,
		}
		// Configure Base URL
		var err error
		openshift.BaseURL, err = url.Parse(convert.Shifter.ClusterConfig.BaseUrl)
		if err != nil {
			panic(err)
		}

		// Confirm Project/Namespace Exists
		_, err = openshift.Apis.Project.Get(item.Namespace.ObjectMeta.Name)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		// Confirm Project/Namespace Exists
		deploymentConfig, err := openshift.Apis.DeploymentConfig.Get(item.Namespace.ObjectMeta.Name, item.DeploymentConfig.ObjectMeta.Name)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		u, err := json.Marshal(deploymentConfig)
		if err != nil {
			panic(err)
		}
		convertedObject := processor.Processor(u, "DeploymentConfig", nil)
		fmt.Println(reflect.TypeOf(convertedObject))

		count = (int(idx) + 1)
	}

	// Construct API Endpoint Response
	r := ResponseConvert{
		UUID:    uuid,
		Message: "Converted..." + string(count) + " Objects",
	}
	ctx.JSON(http.StatusOK, r)
}
