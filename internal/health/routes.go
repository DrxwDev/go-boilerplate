package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Register(rg *gin.RouterGroup)
}

type routes struct{}

func NewRoutes() Routes {
	return &routes{}
}

func (r *routes) Register(rg *gin.RouterGroup) {
	rg.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})
}
