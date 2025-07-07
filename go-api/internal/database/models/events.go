package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	ID          int       `json:"envet_id"`
	Owner       int       `json:"owner_id" binding:"required"`
	Name        string    `json:"event_name" binding:"required,min=3"`
	Description string    `json:"event_description" binding:"required,min=10"`
	Date        time.Time `json:"event_date" binding:"required,datetime=2004-09-11"`
	Location    string    `json:"location" binding:"required,min=5"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id,name,description,date,location) VALUES ($1,$2,$3,$4,$5)"

	return m.DB.QueryRowContext(ctx, query, event.Owner, event.Name, event.Description, event.Date, event.Location).Scan(&event)
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events"

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Owner, &event.Name, &event.Description, &event.Date, &event.Location)

		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events where id = $1"

	var event Event
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&event.ID, &event.Owner, &event.Name, &event.Description, &event.Date, &event.Location)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil ,err
	}


	return &event, nil
}
