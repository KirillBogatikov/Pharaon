package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"net/http"
	"pharaon-auth/data"
	"pharaon-auth/service"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	c := &data.Credentials{}

	err := server.ReadJson(r, c)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, result, err := service.Validate(c)
	if err != nil {
		server.InternalServerError(w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(result, w)
}
