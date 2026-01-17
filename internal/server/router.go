package server

import (
	"github.com/DrxwDev/boilerplate/internal/config"
	"github.com/gin-gonic/gin"
)

func NewGinEngine(srv *config.AppConfig) *gin.Engine {
	// gin.SetMode(os.Getenv("GIN_MODE")) Production Only
	router := gin.Default()
	return router
}
