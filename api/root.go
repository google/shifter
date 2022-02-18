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
	"github.com/gin-gonic/gin"
)

// Instanciate gin-gonic HTTP Server
func InitServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()
	return server, nil

}

// Setup gin-gonic HTTP Server Routes
func (server *Server) setupRouter() {

	// Create Default Gin Router
	router := gin.Default()
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
			s.GET("/healthz", Healthz) // Operations Health Check
			s.GET("/settings", Settings) // Server Settings
		}
	}

	// Setup Server Router for Gin Server Instance.
	server.router = router
}

// Start gin-gonic HTTP Server on Specific Address
func (server *Server) Start(serverAddress string, serverPort string) error {
	return server.router.Run(serverAddress + ":" + serverPort)
}

// Standard API Error Response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

