package server

import (
	"github.com/DrxwDev/boilerplate/internal/health"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, health health.Routes) {
	apiV1 := router.Group("/api/v1")

	health.Register(apiV1)
}
