package main

import (
	"BoxApi/config"
	"BoxApi/server"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	var configFile = flag.String("config", "config/config.json", "set configuration json-file (some_file.json)")
	flag.Parse()

	config := config.NewConfig(*configFile)
	f, err := os.OpenFile(config["PathJournal"], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	var chAbort = make(chan os.Signal)
	signal.Notify(chAbort, os.Interrupt, syscall.SIGUSR1)
	go func() {
		<-chAbort
		func() {
			log.Println("Stopped Service")
		}()
		os.Exit(1)
	}()

	log.SetOutput(f)
	log.Println("Launch Service")

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
