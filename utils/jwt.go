package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "7Hvrd9J8EvRV8+TIKPdXKMTOI2to5fYq7rEEaxxyUHQ="

// Generate a JSON Web Token (JWT) for a given email and user ID.
func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	// Sign the token with the secret key
	return token.SignedString([]byte(JWT_SECRET))
}

// Verifies a JWT token and returns the user ID if the token is valid.
func VerifyToken(token string) (int64, error) {
	// Parse the JWT token -> use are using an anonymous function to verify the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		// Verify that the token is signed with the correct method
		// The token method should be HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method!")
		}
		// Return the secret key
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token!")
	}

	// Check if the token is valid
	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("Invalid token!")
	}

	// Extracting userID from the token for accessing the protected routes
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token!")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil
}
