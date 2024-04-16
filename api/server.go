package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/matheuspolitano/GoMasterDB/db/sqlc"
)

type Server struct {
	router *gin.Engine
	store  *db.Queries
}

func NewServer(store *db.Queries) *Server {

	server := Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/", server.createAccount)

	server.router = router

	return &server

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) StartServer() error {
	err := s.router.Run(":9090")
	return err
}
