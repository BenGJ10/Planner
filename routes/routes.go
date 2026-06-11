package routes

import (
	"event-manager-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", home)
	server.GET("/health", healthCheck)

	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)

	// Group authenticated routes together
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Add authenticated routes
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
