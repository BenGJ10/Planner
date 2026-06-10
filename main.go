package main

import (
	"event-manager-app/db"
	"event-manager-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitDB()

	// Initialize Gin router
	server := gin.Default()

	server.GET("/", home)

	server.GET("/health", healthCheck)

	server.GET("/events", getAllEvents)

	server.POST("/events", createEvent)

	server.Run(":8080")
}

// Handler for homepage
func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Event Management system!",
	})
}

// Handler for checking health status
func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "up",
		"version": "0.1",
	})
}

// Fetch all events from the database
func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events! Try again.",
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

// Create a new event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}
	// event.ID = 1
	// event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event! Try again.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})

}
