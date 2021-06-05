package main

import (
    "github.com/Projector-Solutions/Pharaon-tools/server"
    "io"
    "log"
    "net/http"
    "pharaon-personal/handler"
    "pharaon-personal/service"
)

type PersonalServer struct {
    *server.RestServer
}

func NewServer() (*PersonalServer, error) {
    restServer, err := server.NewServer()
    if err != nil {
        return nil, err
    }
    
    restServer.Router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
        if _, err := io.WriteString(w, "Personal is alive"); err != nil {
            log.Printf("error to write response: %s", err)
        }
    })
    
    err = service.InitRepository()
    if err != nil {
        return nil, err
    }
    
    r := restServer.Router
    r.HandleFunc("/user/{user_id}", handler.GetHandler).Methods("GET")
    r.HandleFunc("/user", handler.SignupHandler).Methods("POST")
    r.HandleFunc("/user/{user_id}", handler.UpdateHandler).Methods("PUT")
    r.HandleFunc("/user/{user_id}", handler.DeleteHandler).Methods("DELETE")
    r.HandleFunc("/validate", handler.ValidateHandler).Methods("POST")
    
    return &PersonalServer{restServer}, nil
}