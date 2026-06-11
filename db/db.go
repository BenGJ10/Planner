package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error

	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	DB, err = sql.Open("pgx", connectionStr)
	if err != nil {
		panic("Could not connect to PostgreSQL")
	}

	err = DB.Ping()
	if err != nil {
		panic("Could not reach PostgreSQL")
	}

	// How many connections can be opened atmost
	DB.SetMaxOpenConns(10)
	// How many connections to keep opened all time
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table! Try again.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		description TEXT,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table! Try again.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table.")
	}
}
