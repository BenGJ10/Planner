package routes

import (
	"event-manager-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// Fetch all events
func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events! Try again.",
		})
		return
	}

	// Return list of events with 200 OK status
	context.JSON(http.StatusOK, events)
}

// Fetch event by ID
func getEventByID(context *gin.Context) {

	// Parse URL parameter for event ID
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event Id! Try again.",
		})
		return
	}

	// Fetch event from the database
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event! Try again.",
		})
		return
	}

	// Return event with 200 OK status
	context.JSON(http.StatusOK, event)
}

// Create a new event
func createEvent(context *gin.Context) {
	var event models.Event
	// Bind JSON to event struct
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	// Placeholder for userId
	// event.UserId = 1

	// Save event to database
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event! Try again.",
		})
		return
	}

	// Return event with 201 Created status
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {
	// Parse URL parameter for event ID
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event Id! Try again.",
		})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event! Try again.",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event! Try again.",
		})
		return
	}

	// Return event with 201 Created status
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event updated successfully",
		"event":   updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	// Parse URL parameter for event ID
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event Id! Try again.",
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event! Try again.",
		})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event! Try again.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
