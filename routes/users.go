package routes

import (
	"event-manager-app/models"
	"event-manager-app/utils"
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

func login(context *gin.Context) {
	var user models.User

	// Bind JSON to user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	// Validate user data
	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed!",
		})
		return
	}

	// Generate JWT token on successful login
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Authentication failed!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successfull!",
		"token":   token,
	})
}
