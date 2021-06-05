package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-card/model"
	"pharaon-card/service"
)

func AddHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	card := &model.Card{}
	err := server.ReadJson(r, card)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, result, err := service.SaveCard(i.Token, card)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if result != nil && !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(card, w)
}
