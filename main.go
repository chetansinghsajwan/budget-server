package main

import (
	"budget-server/db"
	"budget-server/transaction"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("WARNING: Failed to load local env file, err: %s", err.Error())
	}

	db.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /transaction", transaction.HandleCreateTransaction)
	mux.HandleFunc("GET /transaction/{id}", transaction.HandleGetTransaction)
	mux.HandleFunc("PATCH /transaction/{id}", transaction.HandleUpdateTransaction)
	mux.HandleFunc("DELETE /transaction/{id}", transaction.HandleDeleteTransaction)

	log.Println("Server listening at port 8080...")
	serverError := http.ListenAndServe(":8080", mux)
	log.Fatal(serverError)
}
