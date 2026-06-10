package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "yourSecretKey"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(JWT_SECRET)
}
