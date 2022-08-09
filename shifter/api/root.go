/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package api

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config ServerConfig
}

type ServerConfig struct {
	serverAddress   string
	serverPort      string
	storagePlatform string
	//gcsBucket       string
	serverStorage ServerStorage
}

type ServerStorage struct {
	description string
	storageType string
	sourcePath  string
	outputPath  string
}

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

	// Setup & Configure CORS
	config := cors.DefaultConfig()
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"}
	config.AddAllowMethods("OPTIONS", "GET", "POST")
	config.AddAllowHeaders("Origin", "Accept", "Accept-Encoding", "Cache-Control", "X-Requested-With", "X-Custom-Header", "Content-Type", "Content-Length", "Authorization")
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	// Appply CORS
	router.Use(cors.New(config))

	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"OPTIONS", "GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,*/
	/*AllowOriginFunc: func(origin string) bool {
	    return origin == "https://example.com"
	},*/
	/*MaxAge: 12 * time.Hour,
	}))*/

	// Declare API V1 Route Group and Routes
	v1 := router.Group("/api/v1")
	{
		// Convert V1 API Endpoints
		o := v1.Group("/openshift")
		{
			op := o.Group("/projects")
			{
				op.POST("/", server.SOSGetProjects)
				op.POST("/:projectName", server.SOSGetProject)
			}

			dc := o.Group("/deploymentconfigs")
			{
				dc.POST("/", server.SOSGetDeploymentConfigs)
				dc.POST("/:projectName", server.SOSGetDeploymentConfigsByProject)
				dc.POST("/:projectName/:deploymentConfigName", server.SOSGetDeploymentConfig)
			}

		}

		// Convert V1 API Endpoints
		s := v1.Group("/shifter")
		{
			sc := s.Group("/convert")
			{
				sc.POST("/", server.Convert)
			}

			sd := s.Group("/downloads")
			{
				sd.GET("/", server.Downloads)
				sd.GET("/:downloadId", server.Download)
				sd.GET("/:downloadId/file", server.DownloadFile)
			}
		}

		// Status V1 API Endpoints
		st := v1.Group("/status")
		{
			st.GET("/healthz", server.Healthz)   // Operations Health Check
			st.GET("/settingz", server.Settingz) // Server Settingz
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
