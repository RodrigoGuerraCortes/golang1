package main

import (
	"core/internal/api"
	"core/internal/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	db.InitDB()
	defer func() {
		if db.DB != nil {
			db.DB.Close() // Safely close the database connection
		}
	}()

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	api.RegisterRoutes(router)

	// Start the server
	log.Println("Server running on port 8080...")

	router.Run(":8080")

}
