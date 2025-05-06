package main

import (
	"log"
	"net/http"
	"budget-server/transaction"
)

func main() {

	http.HandleFunc("/transaction", transaction.HandleTransaction)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
