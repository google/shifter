package api

// Custom Shifter Server Configuration
type ServerConfig struct {
	serverAddress   string
	serverPort      string
	storagePlatform string
	//gcsBucket       string
	serverStorage ServerStorage
}
