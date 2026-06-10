package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", home)

	server.GET("/health", healthCheck)

	server.GET("/events", getAllEvents)

	server.GET("/events/:id", getEventByID)

	server.POST("/events", createEvent)
}
