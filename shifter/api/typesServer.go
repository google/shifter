package api

import "github.com/gin-gonic/gin"

// HTTP Server Based on gin-gonic
type Server struct {
	router *gin.Engine
	config ServerConfig
}
