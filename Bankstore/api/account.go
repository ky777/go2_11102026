package api

import (
	db "Bankstore/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createAccountRequest struct {
	Owner    string      `json:"owner" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	// десериализация входящего JSON' а
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespons(err))
	}
	// на основе входящего CreateAccountRequest создаем CreateAccountParams
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	// на основе arg создаем новый account
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespons(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type GetAccountRequest struct {
	// URL : /account/1
	// use path(uri) paraneters
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespons(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorRespons(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorRespons(err))
	}
	ctx.JSON(http.StatusOK, account)
}

type ListAccountsRequest struct {
	// URL: /accounts/?page_id=1&page_size=10
	// use query parameters
	PageID   int32 `form:"page_id" building:"required,min=1`
	PageSize int32 `form:"page_size" building:"required,min=5,max=15`
}

func (server *Server) ListAccounts(ctx *gin.Context) {
	var req ListAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errorRespons(err))
		return
	}
	//make query to db
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespons(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
