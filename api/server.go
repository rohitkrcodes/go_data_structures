package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitkrcodes/go_data_structures/lists"
)

type Server struct {
	db     *lists.DoublyList
	router *gin.Engine
}

func NewServer(db *lists.DoublyList) *Server {
	server := &Server{db: db}
	router := gin.Default()

	router.POST("/data", server.addData)
	router.GET("/data", server.listData)
	router.GET("/data/:id", server.getData)
	router.DELETE("/data/:id", server.removeData)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
