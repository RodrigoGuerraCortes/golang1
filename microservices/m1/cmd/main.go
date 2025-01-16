package main

import (
	"log"
	"microservices/m1/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)

	log.Println("M1 service running on port 8081...")
	router.Run(":8081")
}
