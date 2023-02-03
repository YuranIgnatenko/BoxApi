package main

import (
	"BoxApi/config"
	"BoxApi/server"
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	configFile := flag.String("config", "config/config.json", "set configuration json-file (some_file.json)")
	flag.Parse()

	config := config.NewConfig(*configFile)

	server := server.NewServer(config)
	router := gin.Default()
	router.Handle("GET", "/home", server.HandleHomeGet)
	router.Handle("POST", "/home", server.HandleHomePost)
	router.Handle("GET", "/time", server.HandleTimeGet)
	router.Handle("POST", "/time", server.HandleTimePost)
	router.Handle("GET", "/calc", server.HandleCalcGet)
	router.Handle("POST", "/calc", server.HandleCalcPost)
	router.Run(server.Addr())
}
