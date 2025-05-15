package handlers

import (
	"budget-server/db"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionCreate struct {
	Title    string
	OwnerId  uint
	Amount   uint64
	IsCredit bool
	Time     *time.Time
	Tags     []string
}

func HandleCreateTransaction(ctx *gin.Context) {

	var data db.TransactionCreate
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})

		return
	}

	id, err := db.CreateTransaction(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func HandleDeleteTransaction(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err = db.DeleteTransaction(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.Status(http.StatusOK)
}

func HandleUpdateTransaction(ctx *gin.Context) {
	log.Println("Update transaction request received.")
}

func HandleGetTransaction(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var transaction *db.Transaction
	transaction, err = db.GetTransaction(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Transaction": transaction,
	})
}
