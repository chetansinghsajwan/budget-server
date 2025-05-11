package db

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {

	log.Print("Connecting to database...")

	var connString = os.Getenv("DB_CONNECT_STRING")
	if connString == "" {
		log.Fatal("Connecting to database failed, empty connection string")
		return
	}

	var err error
	DB, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatalf("Connecting to database failed, err: %s", err.Error())
		return
	}

	log.Print("Connecting to database done.")
}
