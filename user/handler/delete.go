package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/Projector-Solutions/Pharaon-tools/uuid"
	"log"
	"net/http"
	"pharaon-user/service"
)

func DeleteHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	if !uuid.Equals(id, i.CredentialsId) {
		//todo: UAC
	}

	found, err := service.Delete(id)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(nil, w)
}
