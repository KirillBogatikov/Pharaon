package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"net/http"
	"pharaon-auth/service"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	c, err := service.GetCredentials(userId)
	if err == service.NotFoundError {
		server.NotFound(w)
		return
	}

	if err != nil {
		server.InternalServerError(w)
		return
	}

	server.Ok(c, w)
}
