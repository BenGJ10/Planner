package main

import (
	"event-manager-app/db"
	"event-manager-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Could not load .env")
	}

	// Initialize database connection
	db.InitDB()

	// Initialize Gin router
	server := gin.Default()

	// Register all the endpoints
	routes.RegisterRoutes(server)

	// Start the server
	server.Run(":8080")

}
