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
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseStatusSettings struct {
	Timestamp          string `json:"timestamp"`
	RunningPort        string `json:"runningPort"`
	RunningHost        string `json:"runningHost"`
	StorageType        string `json:"storageType"`
	StorageDescription string `json:"storageDescription"`
	StorageSourcePath  string `json:"storageSourcePath"`
	StorageOutputPath  string `json:"storageOutputPath"`
	Version            int    `json:"version"`
	Status             int    `json:"status"`
	Message            string `json:"message"`
}

type ResponseStatusHealthz struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

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

	ctx.JSON(http.StatusOK, r)
}

func (server *Server) Healthz(ctx *gin.Context) {
	r := ResponseStatusHealthz{}
	r.Timestamp = ""
	r.Status = http.StatusOK
	r.Message = "Shifter Server is reachable."
	ctx.JSON(http.StatusOK, r)
}
