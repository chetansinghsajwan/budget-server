package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Phone     string `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Secret struct {
	ID        uint   `gorm:"primaryKey"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	UserID    uint   `gorm:"not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

var DB *gorm.DB

func DeleteUserByKey(id uint64) error {
	return DB.Delete(&User{}, id).Error
}

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

	log.Print("Migrating schema...")

	// err = DB.AutoMigrate(
	// 	&User{},
	// 	&Secret{},
	// 	&Transaction{},
	// )

	if err != nil {
		log.Fatalf("Migrating schema failed, err: %s", err)
	}

	log.Print("Migrating schema done.")
}
