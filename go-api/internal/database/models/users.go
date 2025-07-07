package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"user_id"`
	Name     string `json:"user_name" binding:"required,min=3"`
	Email    string `json:"user_email" binding:"required,email"`
	Password string `json:"user_passwd" binding:"required" `
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"

	return m.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.Id)

}
