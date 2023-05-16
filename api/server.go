package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	db "github.com/santhoshvempali/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {

	server := &Server{
		store: store,
	}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		fmt.Println(ok)
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.POST("/transfer", server.CreateTransfer)
	router.POST("/user", server.createUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
