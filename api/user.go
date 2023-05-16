package api

import (
	"net/http"

	db "github.com/santhoshvempali/simplebank/db/sqlc"
	"github.com/santhoshvempali/simplebank/util"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}
	ctx.JSON(http.StatusOK, user)

}

// type getAccountRequest struct {
// 	ID int64 `uri:"id"`
// }

// func (server *Server) getAccount(ctx *gin.Context) {
// 	var req getAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errResponse(err))
// 		return
// 	}
// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, account)
// }
