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
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	// Use prepared statement to prevent SQL injection
	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Use `defer` to ensure the statement is closed after the function returns
	defer prepared_stmt.Close()

	// Use the hashed password instead of plain-text password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := prepared_stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId

	return err
}

func (u *User) Validate() error {
	query := "select id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials! Try again.")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("Invalid credentials! Try again.")
	}

	return nil
}
