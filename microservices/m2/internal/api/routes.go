package api

import (
	"microservices/m2/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/m2", handlers.M2Handler)
	}
}
