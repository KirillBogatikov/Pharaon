package main

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"io"
	"log"
	"net/http"
	"pharaon-migration/handler"
	"pharaon-migration/service"
)

type MigrationServer struct {
	*server.RestServer
}

func NewServer() (*MigrationServer, error) {
	restServer, err := server.NewServer()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Migration is alive"); err != nil {
			log.Printf("error to write response: %s", err)
		}
	})

	err = service.InitRepository()
	if err != nil {
		return nil, err
	}

	err = service.InitService()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/service", handler.RegisterServiceHandler).Methods("POST")

	restServer.Router.HandleFunc("/service/{service}/apply", handler.ApplyMigrationHandler).Methods("POST")

	restServer.Router.HandleFunc("/service/{service}/rollback", handler.RollbackMigrationHandler).Methods("POST")

	return &MigrationServer{restServer}, nil
}
