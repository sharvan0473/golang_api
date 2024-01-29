package api

import (
	"github.com/gin-gonic/gin"
	dbs "sim_bank/db/sqlc"
)

type Server struct {
	store  *dbs.Store
	router *gin.Engine
}

func NewServer(store *dbs.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	server.router = router

	return server
}

func (server Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
