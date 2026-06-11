package middlewares

import (
	"event-manager-app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	header := context.Request.Header.Get("Authorization")

	if header == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authoriaztion failed!",
		})
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authoriaztion failed!",
		})
		return
	}

	context.Set("userID", userID)

	context.Next()
}
