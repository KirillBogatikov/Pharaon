package main

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"io"
	"log"
	"net/http"
	"pharaon-user/handler"
	"pharaon-user/service"
)

type UserServer struct {
	*server.RestServer
}

func NewServer() (*UserServer, error) {
	restServer, err := server.NewServer()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "User is alive"); err != nil {
			log.Printf("error to write response: %s", err)
		}
	})

	err = service.InitRepository()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/validate", handler.ValidateHandler).Methods("POST")
	restServer.Router.HandleFunc("/signup", handler.SignupHandler).Methods("POST")
	restServer.Router.Handle("/check", security.NewSecureHandler(handler.CheckHandler)).Methods("POST")
	restServer.Router.Handle("/{user_id}", security.NewSecureHandler(handler.GetHandler)).Methods("GET")
	restServer.Router.Handle("/{user_id}", security.NewSecureHandler(handler.UpdateHandler)).Methods("PUT")
	restServer.Router.Handle("/{user_id}", security.NewSecureHandler(handler.DeleteHandler)).Methods("DELETE")

	return &UserServer{restServer}, nil
}
