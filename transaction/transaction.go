package transaction

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
    log.Println("Create transaction request received.")
}

func HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
    log.Println("Delete transaction request received.")
}

func HandleUpdateTransaction(w http.ResponseWriter, r *http.Request) {
    log.Println("Update transaction request received.")
}

func HandleGetTransaction(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    log.Printf("Get transaction request received for id '%d'", id)

    w.Header().Set("content-type", "application/json")
    w.Write([]byte(fmt.Sprintf("Get transaction request received for id '%d'", id)))
}

func HandleUnknownTransaction(w http.ResponseWriter, r *http.Request) {
    log.Println("Unknown transaction request received.")
}

func HandleTransaction(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        HandleGetTransaction(w, r)
    case http.MethodPut:
        HandleCreateTransaction(w, r)
    case http.MethodDelete:
        HandleDeleteTransaction(w, r)
    case http.MethodPatch:
        HandleUpdateTransaction(w, r)
    default:
        HandleUnknownTransaction(w, r)
    }
}
