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
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Instanciate gin-gonic HTTP Server
func InitServer(serverAddress string, serverPort string, gcsBucket string) (*Server, error) {
	server := &Server{}

	// Set Server Configuration
	server.config.serverAddress = serverAddress
	server.config.serverPort 	= serverPort
	
	// Configure Server Routes
	server.setupRouter()
	// Configure Server Storage
	server.setupStorage(gcsBucket)

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
		c := v1.Group("/convert")
		{
			c.POST("/yaml/yaml", Yaml2Yaml)
		}

		// Download V1 API Endpoints
		d := v1.Group("/download")
		{
			d.GET("/:uuid/:filename", ConvertedFile) // Download Single Converted File
			d.GET("/:uuid/", ConvertedFilesArchive)  // Download All Converted Files (Archive)
		}

		// Status V1 API Endpoints
		s := v1.Group("/status")
		{
			s.GET("/healthz", Healthz)   // Operations Health Check
			s.GET("/settings", Settings) // Server Settings
		}
	}

	// Setup Server Router for Gin Server Instance.
	server.router = router
}

// Setup Server Storage Data
func (server *Server) setupStorage(gcsBucket string) {
	// Validate Storage Bucket:
	server.config.gcsBucket = gcsBucket
	if gcsBucket != "" {
		// Using GCP Cloud Storage
		fmt.Println("Storage: Using GCP Cloud Storage Bucket")
		server.config.storagePlatform = "GCS"
		server.config.gcsBucket = gcsBucket
		/*
		TODO
		- Add Tests for Access and Permissions on GCS Bucket
		*/
	} else {
		// Using Local Storage
		fmt.Println("Storage: Using Local Disk Storage")
		server.config.storagePlatform = "LOCAL"
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

