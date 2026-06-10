package routes

import (
	"event-manager-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Bind JSON to user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	// Save event to database
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user! Try again.",
		})
		return
	}

	// Return event with 201 Created status
	context.JSON(http.StatusCreated, gin.H{
		"message":    "User created successfully",
		"user_id":    user.ID,
		"user_email": user.Email,
	})

}
