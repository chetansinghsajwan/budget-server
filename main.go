package main

import (
	"log"
	"net/http"
)

func handleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Create transaction request recieved.")
}

func handleDeleteTransasction(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete transaction request recieved.")
}

func handleUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Update transaction request recieved.")
}

func handleGetTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Get transaction request recieved.")
}

func handleUnknwonTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Unknown transaction request recieved.")
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetTransaction(w, r)
	case http.MethodPut:
		handleCreateTransaction(w, r)
	case http.MethodDelete:
		handleDeleteTransasction(w, r)
	case http.MethodPatch:
		handleUpdateTransaction(w, r)
	default:
		handleUnknwonTransaction(w, r)
	}
}

func main() {

	http.HandleFunc("/transaction", handleTransaction)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
