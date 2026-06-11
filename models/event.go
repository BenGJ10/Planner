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
	UserID      int64
}

var events = []Event{}

// Method to save event to the database
func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, location, date_time, description, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	err := db.DB.QueryRow(
		query,
		e.Name,
		e.Location,
		e.DateTime,
		e.Description,
		e.UserID,
	).Scan(&e.ID)

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
		err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"

	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = $1, location = $2, date_time = $3, description = $4
	WHERE id = $5`

	_, err := db.DB.Exec(
		query,
		event.Name,
		event.Location,
		event.DateTime,
		event.Description,
		event.ID,
	)

	return err
}

func (event *Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = $1"

	_, err := db.DB.Exec(query, event.ID)

	return err
}

func (event *Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	prepared_stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = prepared_stmt.Exec(event.ID, userID)
	return err
}
