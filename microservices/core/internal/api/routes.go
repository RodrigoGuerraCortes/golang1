package api

import (
	"core/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/aggregate", handlers.AggregateDataHandler)
	}
}
