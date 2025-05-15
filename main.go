package main

import (
	"budget-server/db"
	"budget-server/handlers"
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

	router.POST("/transaction", handlers.HandleCreateTransaction)
	router.GET("/transaction/:id", handlers.HandleGetTransaction)
	router.PATCH("/transaction/:id", handlers.HandleUpdateTransaction)
	router.DELETE("/transaction/:id", handlers.HandleDeleteTransaction)

	err = router.Run(":8080")
	log.Fatal(err)
}
