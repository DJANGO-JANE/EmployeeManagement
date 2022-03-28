package api

import (

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	router *gin.Engine
	DB *sqlx.DB
}

/*func NewServer(DB *sqlx.DB) *Server{
	server := &Server{DB: DB}
	router := gin.Default()

	router.GET("/Employees",server. )
	server.router = router
	return server
}*/