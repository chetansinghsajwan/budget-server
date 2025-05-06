package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	log.Print("Connecting to database...")

	connString := os.Getenv("DB_CONNECT_STRING")
	if connString == "" {
		log.Fatal("Connecting to database failed, empty connection string")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Connecting to database failed, err: %s", err.Error())
		return
	}

	log.Print("Connecting to database done.")
}
