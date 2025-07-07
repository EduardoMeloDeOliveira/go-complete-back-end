package main

import (
	"database/sql"
	"log"
)


func main (){

	dbPathConnection := "postgres://root:1033@localhost:5432/goDB?sslmode=disable"

	db, err := sql.Open("postgres", dbPathConnection); if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}