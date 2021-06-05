package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-card/service"
)

func GetHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	id, err := server.ReadPathUUID(CardIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	card, err := service.GetCard(i.Token, id)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if card == nil {
		server.NotFound(w)
		return
	}

	server.Ok(card, w)
}
