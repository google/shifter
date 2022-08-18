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
	"errors"
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shifter/lib"
)

type Server struct {
	router *gin.Engine
	config ServerConfig
}

type ServerConfig struct {
	serverAddress   string
	serverPort      string
	storagePlatform string
	//gcsBucket       string //TODO ? What is this?
	serverStorage ServerStorage
}

type ServerStorage struct {
	description string
	storageType string
	sourcePath  string
	outputPath  string
}

func InitServer(serverAddress string, serverPort string, sourcePath string, outputPath string) (*Server, error) {
	// Instantiate Shifter Server Struct
	server := &Server{}

	// Set Server Configuration Elements
	// TODO - Need to tidy up and ensure default handling with constants
	server.config.serverAddress = serverAddress // Set HTTP Server Address
	server.config.serverPort = serverPort       // Set HTTP Server Port

	// Configure Server Routes
	err := server.setupServer()
	if err != nil {
		log.Printf("❌  Error: Failed to Setup Shifter Server (gin-gonic) Routes, Unable to continue.")
		return server, err
	}

	err = server.setupStorage(sourcePath, outputPath)
	if err != nil {
		// Error Setting Shifter Server (gin-gonic) Routes
		log.Printf("❌  Error: Failed to Setup Shifter Server Storage, Unable to continue.")
		return server, err
	}

	return server, nil
}

// Setup gin-gonic HTTP Server Routes
func (server *Server) setupServer() error {
	// TODO - Set Wrapper and CLI Flag for "DebugMode" and inlcude optional this in that flag
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default() // TODO - Explore Non-Default options for Logging.

	// Setup & Configure Gin-Gonic CORS
	config := cors.DefaultConfig()
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"} // TODO Enable the ability to limit this with CLI for Security
	config.AddAllowMethods("OPTIONS", "GET", "POST")
	config.AddAllowHeaders("Origin", "Accept", "Accept-Encoding", "Cache-Control", "X-Requested-With", "X-Custom-Header", "Content-Type", "Content-Length", "Authorization")
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// Declare API V1 Route Group and Routes
	log.Printf("💡 INFO: API Declare Route Group and Routes: /api/v1")
	v1 := router.Group("/api/v1")
	{
		log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift")
		o := v1.Group("/openshift")
		{
			log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift/projects")
			op := o.Group("/projects")
			{
				op.POST("/", server.SOSGetProjects)
				op.POST("/:projectName", server.SOSGetProject)
			}

			log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift/deploymentconfigs")
			dc := o.Group("/deploymentconfigs")
			{
				dc.POST("/", server.SOSGetDeploymentConfigs)
				dc.POST("/:projectName", server.SOSGetDeploymentConfigsByProject)
				dc.POST("/:projectName/:deploymentConfigName", server.SOSGetDeploymentConfig)
			}

		}

		// Convert Shifter API Endpoints
		log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter")
		s := v1.Group("/shifter")
		{
			log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter/convert")
			sc := s.Group("/convert")
			{
				sc.POST("/", server.Convert)
			}

			log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter/downloads")
			sd := s.Group("/downloads")
			{
				sd.GET("/", server.Downloads)
				sd.GET("/:downloadId", server.Download)
				sd.GET("/:downloadId/file", server.DownloadFile)
			}
		}

		// Convert Shifter Server Status API Endpoints
		log.Printf("💡 INFO: API Declare Route Group and Routes: /status")
		st := v1.Group("/status")
		{
			st.GET("/healthz", server.Healthz)
			st.GET("/settingz", server.Settingz)
		}
	}

	server.router = router

	// Setup Successful
	log.Printf("✅ SUCCESS: Shifter Storage Settings Configured")
	return nil
}

// Setup Server Storage Data
func (server *Server) setupStorage(sourcePath string, outputPath string) error {

	// TODO - Rework this Function to look at CLI Input and Defaults and Bolster this with Validations
	server.config.serverStorage.sourcePath = "data/source"
	server.config.serverStorage.outputPath = "data/output"
	server.config.serverStorage.description = "Shifter Server is Connected to Local Storage"

	// Check if SourcePath is Configured for GCS
	if strings.Contains(sourcePath, "gs://") {
		log.Printf("💡 INFO: Configuring Shifter to utilize GCP Cloud Storage Buckets")
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
	} else {
		// Not Using GCS. Use Local
		log.Printf("💡 INFO: Configuring Shifter to utilize GCP Cloud Storage Buckets")
	}

	// Output all Storage Settings
	log.Printf("🔧 DEBUG: [Shifter Storage Settings] - [Storage Type: %s]", server.config.serverStorage.storageType)
	log.Printf("🔧 DEBUG: [Shifter Storage Settings] - [Description: %s]", server.config.serverStorage.description)
	log.Printf("🔧 DEBUG: [Shifter Storage Settings] - [Source Path: %s]", server.config.serverStorage.sourcePath)
	log.Printf("🔧 DEBUG: [Shifter Storage Settings] - [Output Path: %s]", server.config.serverStorage.outputPath)

	return nil
}

func (server *Server) Start() error {
	lib.CLog("debug", "Shifter server starting")

	err := server.router.Run(server.config.serverAddress + ":" + server.config.serverPort)
	if err != nil {
		log.Printf("🌐 ❌ ERROR: Failed to Run Shifter Server, Unable to continue.")
		return err
	}
	lib.CLog("info", "Shifter server listening at "+server.config.serverAddress+":"+server.config.serverPort)

	return nil
}

// Standard API Error Response
func errorResponse(err error) gin.H {
	// Log the High Level Error
	log.Printf("🌐 ❌ ERROR [API]: %s", err)
	// Return Error that will be passed on to Client
	return gin.H{"error": err.Error()}
}
