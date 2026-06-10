package models

import "event-manager-app/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Method to save event to the database
func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	// Use prepared statement to prevent SQL injection
	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Use `defer` to ensure the statement is closed after the function returns
	defer prepared_stmt.Close()

	result, err := prepared_stmt.Exec(u.Email, u.Password)
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
