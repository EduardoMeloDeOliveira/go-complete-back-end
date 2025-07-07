package database

import "database/sql"

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"user_id"`
	Name     string `json:"user_name" binding:"required,min=3"`
	Email    string `json:"user_email" binding:"required,email"`
	Password string `json:"user_passwd" binding:"required" `
}
