package main

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	pconfig "github.com/Projector-Solutions/Pharaon-config/proxy"
	"log"
	"net/http"
	"os"
	"os/signal"
	log2 "pharaon-proxy/log"
	"pharaon-proxy/settings"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	logger, err := log2.NewLogger()
	if err != nil {
		panic(err)
	}

	set, err := settings.Load()
	if err != nil {
		panic(err)
	}

	proxy := NewProxy(logger, set)

	for _, service := range config.Global.Services {
		if service.Name == "proxy" {
			continue
		}

		proxy.Register(service.Name, service.HttpConfig.ApiUrl)
	}

	go func() {
		log.Printf("Starting proxy server on %s\n", pconfig.Http.BindAddress)
		if err := http.ListenAndServe(pconfig.Http.BindAddress, proxy); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	log.Println("Server gracefully stopped")
}
