package api

import (
	"microservices/m1/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/m1", handlers.M1Handler)
	}
}
