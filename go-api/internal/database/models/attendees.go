package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID    int `json: "attendee_id"`
	User  int `json : "user_id"`
	Event int `json : "event_id"`
}
