package main

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"io"
	"log"
	"net/http"
	"pharaon-card/handler"
	"pharaon-card/service"
)

type CardServer struct {
	*server.RestServer
}

func NewServer() (*CardServer, error) {
	restServer, err := server.NewServer()
	if err != nil {
		return nil, err
	}

	restServer.Router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Card is alive"); err != nil {
			log.Printf("error to write response: %s", err)
		}
	})

	err = service.InitRepository()
	if err != nil {
		return nil, err
	}

	restServer.Router.Handle("/list", security.NewSecureHandler(handler.ListHandler)).Methods("POST")
	restServer.Router.Handle("/{card_id}", security.NewSecureHandler(handler.GetHandler)).Methods("GET")
	restServer.Router.Handle("/", security.NewSecureHandler(handler.AddHandler)).Methods("POST")
	restServer.Router.Handle("/{card_id}", security.NewSecureHandler(handler.SaveHandler)).Methods("PUT")
	restServer.Router.Handle("/{card_id}", security.NewSecureHandler(handler.DeleteHandler)).Methods("DELETE")

	restServer.Router.Handle("/tag/list", security.NewSecureHandler(handler.TagListHandler)).Methods("GET")
	restServer.Router.Handle("/tag/search", security.NewSecureHandler(handler.TagAutocompleteHandler)).Methods("GET")
	restServer.Router.Handle("/tag", security.NewSecureHandler(handler.TagCreateHandler)).Methods("POST")
	restServer.Router.Handle("/tag/{tag_id}", security.NewSecureHandler(handler.TagDeleteHandler)).Methods("DELETE")

	return &CardServer{restServer}, nil
}
