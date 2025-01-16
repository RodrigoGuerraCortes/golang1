package main

import (
	"log"
	"microservices/m2/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)

	log.Println("M2 service running on port 8082...")
	router.Run(":8082")
}
