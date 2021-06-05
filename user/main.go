package main

import (
	config "github.com/Projector-Solutions/Pharaon-config/user"
	"log"
	"os"
	"os/signal"
	"pharaon-user/service"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	log.Println("* * * * * MIGRATIONS * * * * *")
	err := service.Migrate()
	if err != nil {
		panic(err)
	}

	log.Println("* * * * * WEB SERVER * * * * *")
	server, err := NewServer()
	if err != nil {
		panic(err)
	}

	server.Start(config.Service.HttpConfig.BindAddress)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	log.Println("Server gracefully stopped")
}
