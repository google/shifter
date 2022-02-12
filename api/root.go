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
	"shifter/api/download"
	"shifter/api/convert" 

	"github.com/gin-gonic/gin"
)

// Instanciate gin-gonic HTTP Server
func initServer() (*Server, error) {
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
		v1.POST("/convert/yaml/yaml", convert.Yaml2Yaml)
		// Download V1 API Endpoints
		v1.GET("/download/:uuid/:filename", download.convertedFile)
	}

	// Setup Server Router for Gin Server Instance.
	server.router = router
}

// Start gin-gonic HTTP Server on Specific Address
func (server *Server) ServerStart(address string) error {
	return server.router.Run(address)
}

// Standard API Error Response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

/*
func Server(httpPort string, flags map[string]string) {

	// Production Mode
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/convert/yaml/yaml", convert.Ep_ConvertYaml2Yaml)
		v1.GET("/download/:uuid/:filename", Download)
	}
	r.Run((":" + httpPort)) // listen and serve on 0.0.0.0:{httpPort} (for windows "localhost:{httpPort}")
}*/
