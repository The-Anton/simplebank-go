package api

import (
	"net/http"
	db "simplebank-go/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	AccountName string `json:"account_name" binding:"required"`
	Currency    string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		AccountName: req.AccountName,
		Balance:     0,
		Currency:    req.Currency,
	}

	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, acc)
}
