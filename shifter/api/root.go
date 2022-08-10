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
)

// Shifter Server Core Structure
type Server struct {
	router *gin.Engine
	config ServerConfig
}

// Shifter Server Configuration Structure
type ServerConfig struct {
	serverAddress   string
	serverPort      string
	storagePlatform string
	//gcsBucket       string //TODO ? What is this?
	serverStorage ServerStorage
}

// Shifter Server Storage Structure
type ServerStorage struct {
	description string
	storageType string
	sourcePath  string
	outputPath  string
}

// Instantiate gin-gonic HTTP Server
func InitServer(serverAddress string, serverPort string, sourcePath string, outputPath string) (*Server, error) {
	// Instantiate Shifter Server Struct
	server := &Server{}

	// Set Server Configuration Elements
	// TODO - Need to tiday up and ensure default handling with constants
	server.config.serverAddress = serverAddress // Set HTTP Server Address
	server.config.serverPort = serverPort       // Set HTTP Server Port

	// Configure Server Routes
	err := server.setupServer()
	if err != nil {
		// Error Setting Shifter Server (gin-gonic) Routes
		log.Printf("❌  Error: Failed to Setup Shifter Server (gin-gonic) Routes, Unable to continue.")
		return server, err
	}

	// Configure Server Storage
	err = server.setupStorage(sourcePath, outputPath)
	if err != nil {
		// Error Setting Shifter Server (gin-gonic) Routes
		log.Printf("❌  Error: Failed to Setup Shifter Server Storage, Unable to continue.")
		return server, err
	}

	// Success: Return Server Instance
	log.Printf("✅ SUCCESS: Shifter Server Configured")
	return server, nil

}

// Setup gin-gonic HTTP Server Routes
func (server *Server) setupServer() error {
	// TODO - Set Wrapper and CLI Flag for "DebugMode" and inlcude optional this in that flag
	// Set Release Mode
	gin.SetMode(gin.ReleaseMode)

	// Default returns an Gin Engine instance with the Logger and Recovery middleware already attached.
	log.Printf("💡 INFO: Creating Gin-Gonic Engine")
	router := gin.Default() // TODO - Explore Non-Default options for Logging.

	// Setup & Configure Gin-Gonic CORS
	log.Printf("💡 INFO: Setup & Configure Gin-Gonic CORS")
	config := cors.DefaultConfig()
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"} // TODO Enable the ability to limit this with CLI for Security
	config.AddAllowMethods("OPTIONS", "GET", "POST")
	config.AddAllowHeaders("Origin", "Accept", "Accept-Encoding", "Cache-Control", "X-Requested-With", "X-Custom-Header", "Content-Type", "Content-Length", "Authorization")
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	// Add CORS Middlemare to Gin-Gonic EngineUse()
	router.Use(cors.New(config))

	// Declare API V1 Route Group and Routes
	log.Printf("💡 INFO: API Declare Route Group and Routes: /api/v1")
	v1 := router.Group("/api/v1")
	{
		// Convert Openshift API Endpoints
		log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift")
		o := v1.Group("/openshift")
		{
			// API Endpoints - Projects
			log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift/projects")
			op := o.Group("/projects")
			{
				// API Endpoint - Get Projects
				op.POST("/", server.SOSGetProjects)
				// API Endpoint - Get Project [By Name]
				op.POST("/:projectName", server.SOSGetProject)
			}

			// API Endpoints - DeploymentConfigs
			log.Printf("💡 INFO: API Declare Route Group and Routes: /openshift/deploymentconfigs")
			dc := o.Group("/deploymentconfigs")
			{
				// API Endpoint - Get DeploymentConfigs
				dc.POST("/", server.SOSGetDeploymentConfigs)
				// API Endpoint - Get DeploymentConfig by OS ProjectName
				dc.POST("/:projectName", server.SOSGetDeploymentConfigsByProject)
				// API Endpoint - Get DeploymentConfig by OS ProjectName & OS DeploymentConfigName
				dc.POST("/:projectName/:deploymentConfigName", server.SOSGetDeploymentConfig)
			}

		}

		// Convert Shifter API Endpoints
		log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter")
		s := v1.Group("/shifter")
		{
			// API Endpoints - Converter
			log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter/convert")
			sc := s.Group("/convert")
			{
				sc.POST("/", server.Convert)
			}

			// API Endpoints - Downloads
			log.Printf("💡 INFO: API Declare Route Group and Routes: /shifter/downloads")
			sd := s.Group("/downloads")
			{
				// API Endpoint - Get Downloads
				sd.GET("/", server.Downloads)
				// API Endpoint - Get Download by ID
				sd.GET("/:downloadId", server.Download)
				// API Endpoint - Get Download File by ID
				sd.GET("/:downloadId/file", server.DownloadFile)
			}
		}

		// Convert Shifter Server Status API Endpoints
		log.Printf("💡 INFO: API Declare Route Group and Routes: /status")
		st := v1.Group("/status")
		{
			// API Endpoint - Operational Shifter Server Health Checks
			st.GET("/healthz", server.Healthz)
			// API Endpoint - Operational Shifter Server Settings
			st.GET("/settingz", server.Settingz)
		}
	}

	// Setup Server Router for Gin Server Instance.
	log.Printf("💡 INFO: Attach Gin-Gonic Engine to Shifter Instance")
	server.router = router

	// Setup Successful
	log.Printf("✅ SUCCESS: Shifter Storage Settings Configured")
	return nil
}

// Setup Server Storage Data
func (server *Server) setupStorage(sourcePath string, outputPath string) error {

	// TODO - Rework this Function to look at CLI Input and Defaults and Bolster this with Validations

	// Set Default Storage Type to LCL (Local Storage)
	server.config.serverStorage.storageType = LCL
	// Set Default Storage Source Path to "./data/source"
	server.config.serverStorage.sourcePath = "data/source"
	// Set Default Storage Source Path to "./data/output"
	server.config.serverStorage.outputPath = "data/output"
	// Set Default Storage Description
	server.config.serverStorage.description = "Shifter Server is Connected to Local Storage"

	// Check if SourcePath is Configured for GCS
	if strings.Contains(sourcePath, "gs://") {
		// Using GCP Cloud Storage
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

	// Setup Successful
	log.Printf("✅ SUCCESS: Shifter Gin-Gonic Engine & Routes")
	return nil
}

// Start gin-gonic HTTP Server on Specific Address
func (server *Server) Start() error {
	// Log Server Run Diagnostics
	log.Printf("🌐 💡 INFO: Shifter Server Starting")
	log.Printf("🌐 🔧 DEBUG: [Shifter Server Settings] - [Port: %s]", server.config.serverPort)
	log.Printf("🌐 🔧 DEBUG: [Shifter Server Settings] - [Hostname: %s]", server.config.serverAddress)
	log.Printf("🌐 ✅ SUCCESS: Shifter Server Running")

	// Run Server
	err := server.router.Run(server.config.serverAddress + ":" + server.config.serverPort)
	if err != nil {
		// Error Setting Shifter Server (gin-gonic) Routes
		log.Printf("🌐 ❌ ERROR: Failed to Run Shifter Server, Unable to continue.")
		return err
	} else {
		// Succesfully Started Shifter Server
		log.Printf("🌐 ✅ SUCCESS: Shifter Server Started")
		return nil
	}

	// Unknown Error when Starting Shifter Server. Return Error
	return errors.New("🌐 ❌ ERROR: Unknonw Error starting Shifter Server, Unable to continue.")
}

// Standard API Error Response
func errorResponse(err error) gin.H {
	// Log the High Level Error
	log.Printf("🌐 ❌ ERROR [API]: %s", err)
	// Return Error that will be passed on to Client
	return gin.H{"error": err.Error()}
}
