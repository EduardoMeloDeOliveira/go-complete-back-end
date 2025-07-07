package database

import "database/sql"


type Models struct {
	Users UserModel
	Event EventModel 
	Attendee AttendeeModel
}


func NewModels(db *sql.DB) Models{
	return  Models{
		Users: UserModel{DB : db},
		Event: EventModel{DB: db},
		Attendee:  AttendeeModel{DB: db},
	}
}