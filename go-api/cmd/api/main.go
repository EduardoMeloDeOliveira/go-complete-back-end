package main

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	database "go-rest-api/internal/database/models"
	"go-rest-api/internal/env"
	"log"
)

type Application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {

	dbPathConnection := "postgres://root:1033@localhost:5432/goDB?sslmode=disable"

	db, err := sql.Open("postgres", dbPathConnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &Application{
		port:      env.GetEnvInt("PORT", 8081),
		jwtSecret: env.GetEnvString("JWT_SECRETE", "secret"),
		models:    models,
	}

	if err := app.server(); err != nil {
		log.Fatal(err)
	}
}
