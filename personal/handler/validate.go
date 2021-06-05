package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"net/http"
	"pharaon-personal/model"
	"pharaon-personal/service"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	data := &model.Data{}
	err := server.ReadJson(r, data)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, result, err := service.Validate(data)
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
