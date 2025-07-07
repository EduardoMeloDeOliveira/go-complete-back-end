package database

import (
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	ID          int       `json : "envet_id"`
	Owner       int       `json : "owner_id"`
	Name        string    `json : "event_name"`
	Description string    `json : "event_description"`
	Date        time.Time `json : "event_date"`
	Location    string    `json : "location" `
}
