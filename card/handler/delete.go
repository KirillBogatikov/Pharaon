package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-card/service"
)

func DeleteHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	id, err := server.ReadPathUUID(CardIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, err := service.DeleteCard(id)
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
