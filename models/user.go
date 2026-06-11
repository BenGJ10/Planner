package models

import (
	"errors"
	"event-manager-app/db"
	"event-manager-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Method to save event to the database
func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES ($1, $2)
	RETURNING id`

	// Use the hashed password instead of plain-text password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Execute the query and return the generated ID
	err = db.DB.QueryRow(
		query,
		u.Email,
		hashedPassword,
	).Scan(&u.ID)

	return err
}

func (u *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = $1"

	var retrievedPassword string

	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials! Try again.")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("Invalid credentials! Try again.")
	}

	return nil
}
