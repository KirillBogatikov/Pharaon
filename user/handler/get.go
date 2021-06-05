package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/Projector-Solutions/Pharaon-tools/uuid"
	"log"
	"net/http"
	"pharaon-user/service"
)

func GetHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	if !uuid.Equals(id, i.CredentialsId) {
		//todo: UAC
	}

	user, err := service.GetById(id)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if user == nil {
		server.NotFound(w)
		return
	}

	server.Ok(user, w)
}
