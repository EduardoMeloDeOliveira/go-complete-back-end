package database

import "database/sql"

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json: "user_id"`
	Name     string `json : "user_name"`
	Email    string `json : "user_email"`
	Password string `json : "user_passwd" `
}
