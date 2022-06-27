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
	"net/http"

	"github.com/gin-gonic/gin"
)

const GCS string = "GCS"
const LCL string = "LCL"

func (server *Server) Settingz(ctx *gin.Context) {
	// Construct API Endpoint Response
	r := ResponseStatusSettings{}

	// Server Settings
	r.RunningPort = server.config.serverPort
	r.RunningHost = server.config.serverAddress

	// Storage Settings
	r.StorageType = server.config.serverStorage.storageType
	r.StorageDescription = server.config.serverStorage.description
	r.StorageSourcePath = server.config.serverStorage.sourcePath
	r.StorageOutputPath = server.config.serverStorage.outputPath

	// General Meta Data
	r.Timestamp = ""
	r.Version = 0
	r.Status = http.StatusOK
	r.Message = "Shifter Server Settings."
	// Return JSON API Response
	ctx.JSON(http.StatusOK, r)
}
