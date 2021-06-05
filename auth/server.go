package main

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"io"
	"log"
	"net/http"
	"pharaon-auth/handler"
	"pharaon-auth/service"
)

type AuthService struct {
	*server.RestServer
}

func NewServer() (*AuthService, error) {
	restServer, err := server.NewServer()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Auth is alive"); err != nil {
			log.Printf("error to write response: %s", err)
		}
	})

	err = service.InitRepository()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/auth", handler.AuthHandler).Methods("GET")
	restServer.Router.HandleFunc("/validate", handler.ValidateHandler).Methods("POST")
	restServer.Router.HandleFunc("/user/login", handler.LoginHandler).Methods("GET")
	restServer.Router.HandleFunc("/user/signup", handler.SignupHandler).Methods("POST")
	restServer.Router.HandleFunc("/user/{user_id}", handler.GetHandler).Methods("GET")
	restServer.Router.HandleFunc("/user/{user_id}", handler.UpdateHandler).Methods("PUT")
	restServer.Router.HandleFunc("/user/{user_id}", handler.DeleteHandler).Methods("DELETE")
	restServer.Router.HandleFunc("/history/user/{user_id}", handler.HistoryHandler).Methods("GET")
	restServer.Router.HandleFunc("/history/{history_id}", handler.DeleteHistoryHandler).Methods("DELETE")
	restServer.Router.HandleFunc("/token/user/{user_id}", handler.StartRestore).Methods("POST")
	restServer.Router.HandleFunc("/token/{token}/restore", handler.ApplyRestore).Methods("GET")

	return &AuthService{restServer}, nil
}
