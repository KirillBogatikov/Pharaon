package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func ReadPathUUID(name string, r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	return uuid.Parse(vars[name])
}

func ReadHeaderUUID(name string, h http.Header) (uuid.UUID, error) {
	return uuid.Parse(h.Get(name))
}
