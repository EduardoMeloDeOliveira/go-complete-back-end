package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	dsn := "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir o banco:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Erro ao inicializar driver do postgres:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", 
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Erro ao criar instancia do migrate:", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Erro ao aplicar migrations:", err)
		}
		log.Println("Migrations aplicadas com sucesso.")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Erro ao reverter migrations:", err)
		}
		log.Println("Migrations revertidas com sucesso.")

	default:
		log.Fatal("Direção inválida. Use 'up' ou 'down'.")
	}
}
