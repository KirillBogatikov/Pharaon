package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-user/model"
	"pharaon-user/service"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := server.ReadJson(r, user)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, result, err := service.Validate(user)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(result, w)
}
