package models

import (
	"event-manager-app/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Description string
	UserId      int64
}

var events = []Event{}

// Method to save event to the database
func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, location, datetime, description, user_id)
	VALUES (?, ?, ?, ?, ?)`

	// Use prepared statement to prevent SQL injection
	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Use `defer` to ensure the statement is closed after the function returns
	defer prepared_stmt.Close()

	result, err := prepared_stmt.Exec(e.Name, e.Location, e.DateTime, e.Description, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

// Method to fetch all events from the database
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// Use query instead of prepared statement since no user inputs are involved
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events where id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `
	UPDATE events 
	SET name = ?, location = ?, datetime = ?, description = ?
	WHERE id = ?`

	// Use prepared statement to prevent SQL injection
	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Use `defer` to ensure the statement is closed after the function returns
	defer prepared_stmt.Close()

	_, err = prepared_stmt.Exec(event.Name, event.Location, event.DateTime, event.Description, event.ID)
	return err
}

func (event *Event) DeleteEvent() error {
	query := "DELETE from events WHERE id = ?"

	// Use prepared statement to prevent SQL injection
	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Use `defer` to ensure the statement is closed after the function returns
	defer prepared_stmt.Close()

	_, err = prepared_stmt.Exec(event.ID)
	return err
}
