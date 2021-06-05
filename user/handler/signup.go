package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-user/model"
	"pharaon-user/service"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := server.ReadJson(r, user)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	result, err := service.Signup(user)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	server.Created(user, w)
}
