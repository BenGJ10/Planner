package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err := sql.Open("sqlite3", "event-manager.db")
	if err != nil {
		panic("Could not connect to the database! Try again.")
	}

	// How many connections can be opened atmost
	DB.SetMaxOpenConns(10)
	// How many connections to keep opened all time
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		description TEXT,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table! Try agin.")
	}
}
