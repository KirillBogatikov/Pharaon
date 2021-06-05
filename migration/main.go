package main

import (
	config "github.com/Projector-Solutions/Pharaon-config/migration"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	server, err := NewServer()
	if err != nil {
		panic(err)
	}

	server.Start(config.Service.HttpConfig.BindAddress)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Server gracefully stopped")
}
