package main

import (
	_ "embed"
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	"log"
	"os"
	"os/signal"
	"pharaon-auth/service"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	log.Println("* * * * * MIGRATIONS * * * * *")
	err := service.Migrate()
	if err != nil {
		panic(err)
	}

	log.Println("* * * * * CRON * * * * *")
	err = service.StartAutoClearing()
	if err != nil {
		panic(err)
	}
	log.Println("Cron started. Tokens will be cleared every 12 hours")

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
