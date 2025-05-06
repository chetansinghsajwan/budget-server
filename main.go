package main

import (
	"budget-server/db"
	"budget-server/transaction"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("WARNING: Failed to load local env file, err: %s", err.Error())
	}

	db.Init()

	router := gin.Default()

	router.POST("/transaction", transaction.HandleCreateTransaction)
	router.GET("/transaction/:id", transaction.HandleGetTransaction)
	router.PATCH("/transaction/:id", transaction.HandleUpdateTransaction)
	router.DELETE("/transaction/:id", transaction.HandleDeleteTransaction)

	log.Println("Server listening at port 8080...")
	err = router.Run(":8080")
	log.Fatal(err)
}
