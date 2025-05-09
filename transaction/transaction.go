package transaction

import (
	"budget-server/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleCreateTransaction(ctx *gin.Context) {
	log.Println("Create transaction request received.")
}

func HandleDeleteTransaction(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	log.Printf("Delete request received for id %d", id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_, err = db.DeleteUserByKey(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})

		return
	}

	log.Printf("Delete request completed.")

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

func HandleUnknownTransaction(ctx *gin.Context) {
	log.Println("Unknown transaction request received.")
}
