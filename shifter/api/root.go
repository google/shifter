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
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Instanciate gin-gonic HTTP Server
func InitServer(serverAddress string, serverPort string, sourcePath string, outputPath string) (*Server, error) {
	server := &Server{}

	// Set Server Configuration
	server.config.serverAddress = serverAddress
	server.config.serverPort = serverPort

	// Configure Server Routes
	server.setupRouter()
	// Configure Server Storage
	server.setupStorage(sourcePath, outputPath)

	// Return Server Instance
	return server, nil

}

// Setup gin-gonic HTTP Server Routes
func (server *Server) setupRouter() {

	// Create Default Gin Router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		/*AllowOriginFunc: func(origin string) bool {
		    return origin == "https://example.com"
		},*/
		MaxAge: 12 * time.Hour,
	}))

	// Declare API V1 Route Group and Routes
	v1 := router.Group("/api/v1")
	{
		// Convert V1 API Endpoints
		o := v1.Group("/openshift")
		{
			p := o.Group("/projects")
			{
				p.POST("/", server.OpenShiftGetProjects)
				p.POST("/:projectName", server.OpenShiftGetProject)
			}

		}

		// Convert V1 API Endpoints
		/*c := v1.Group("/convert")
		{
			c.POST("/yaml/yaml", server.Yaml2Yaml)
		}*/

		// Download V1 API Endpoints
		/*d := v1.Group("/download")
		{
			d.GET("/:uuid/:filename", server.ConvertedFile) // Download Single Converted File
			d.GET("/:uuid/", server.ConvertedFilesArchive)  // Download All Converted Files (Archive)
		}*/

		// Status V1 API Endpoints
		s := v1.Group("/status")
		{
			s.GET("/healthz", server.Healthz)   // Operations Health Check
			s.GET("/settingz", server.Settingz) // Server Settingz
		}
	}

	// Setup Server Router for Gin Server Instance.
	server.router = router
}

// Setup Server Storage Data
func (server *Server) setupStorage(sourcePath string, outputPath string) {

	// Set Default Storage Type to LCL (Local Storage)
	server.config.serverStorage.storageType = LCL
	// Set Default Storage Source Path to "./data/source"
	server.config.serverStorage.sourcePath = "data/source"
	// Set Default Storage Source Path to "./data/output"
	server.config.serverStorage.outputPath = "data/output"
	// Set Default Storage Description
	server.config.serverStorage.description = "Shifter Server is Connected to Local Storage"

	if strings.Contains(sourcePath, "gs://") {
		// Using GCP Cloud Storage
		fmt.Println("Storage: Using GCP Cloud Storage Bucket")
		server.config.serverStorage.storageType = GCS
		server.config.serverStorage.description = "Shifter Server is Connected to Google Cloud Storage"

		// If the Provided Bucket for Source and Destination is the Same we will expand the Paths to include Subfolders
		if strings.Compare(sourcePath, outputPath) == 0 {
			sourcePath = sourcePath + "/source"
			outputPath = outputPath + "/output"
		}

		// Set Default Storage Source and Output GCS Bucket
		server.config.serverStorage.sourcePath = sourcePath
		server.config.serverStorage.outputPath = outputPath
	}
}

// Start gin-gonic HTTP Server on Specific Address
func (server *Server) Start() error {
	// Run Server
	fmt.Println(server.config)
	return server.router.Run(server.config.serverAddress + ":" + server.config.serverPort)
}

// Standard API Error Response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
