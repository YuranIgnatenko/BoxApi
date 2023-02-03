package main

import (
	"BoxApi/server"

	"github.com/gin-gonic/gin"
)

func main() {
	server := server.NewServer("localhost", "8080")
	router := gin.Default()
	router.Handle("GET", "/home", server.HandleHomeGet)
	router.Handle("POST", "/home", server.HandleHomePost)
	router.Handle("GET", "/time", server.HandleTimeGet)
	router.Handle("POST", "/time", server.HandleTimePost)
	router.Handle("GET", "/calc", server.HandleCalcGet)
	router.Handle("POST", "/calc", server.HandleCalcPost)
	router.Run(server.Addr())
}
