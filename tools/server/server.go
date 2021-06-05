package server

import (
	"github.com/Projector-Solutions/Pharaon-tools/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RestServer struct {
	Router *mux.Router
}

func NewServer() (*RestServer, error) {
	router := mux.NewRouter()

	router.Use(middleware.LogMiddleware, middleware.JsonTypeMiddleware)

	return &RestServer{router}, nil
}

func (s *RestServer) Start(bindAddress string) {
	go func() {
		http.Handle("/", s.Router)
		if err := http.ListenAndServe(bindAddress, nil); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("HTTP server started on %s\n", bindAddress)
}
