package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/google/uuid"
	"log"
	"net/http"
	"pharaon-card/service"
)

func ListHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	ids := make([]uuid.UUID, 0)
	err := server.ReadJson(r, &ids)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	cards, err := service.GetCards(i.Token, ids...)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	server.Ok(cards, w)
}
