package transaction

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleCreateTransaction(ctx *gin.Context) {
	log.Println("Create transaction request received.")
}

func HandleDeleteTransaction(ctx *gin.Context) {
	log.Println("Delete transaction request received.")
}

func HandleUpdateTransaction(ctx *gin.Context) {
	log.Println("Update transaction request received.")
}

func HandleGetTransaction(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	log.Printf("Get transaction request received for id '%d'", id)
}

func HandleUnknownTransaction(ctx *gin.Context) {
	log.Println("Unknown transaction request received.")
}
