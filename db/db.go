package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DeleteUserByKey(id uint64) (sql.Result, error) {

	var query = fmt.Sprintf(
		`
		delete from users
		where id = %d
		`,
		id,
	)

	return DB.Exec(query)
}

type Transaction struct {
	Id        uint
	Title     string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func GetTransaction(id uint64) (*Transaction, error) {

	var query = fmt.Sprintf(
		`
		select * from transactions
		where id = %d
		`,
		id,
	)

	var row = DB.QueryRow(query)

	var result Transaction
	var err = row.Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

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
