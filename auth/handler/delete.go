package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"net/http"
	"pharaon-auth/service"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	ok, err := service.DeleteAuth(id)
	if err != nil {
		server.InternalServerError(w)
		return
	}

	if ok {
		server.Ok(nil, w)
		return
	}

	server.NotFound(w)
}
