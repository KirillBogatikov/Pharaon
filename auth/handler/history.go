package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/Projector-Solutions/Pharaon-tools/uuid"
	"log"
	"net/http"
	"pharaon-auth/service"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	info := auth(w, r)
	if info == nil {
		return
	}

	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	if !uuid.Equals(id, info.CredentialsId) {
		log.Printf("access denied %s -> %s", info.CredentialsId, id)
		server.Forbidden(w)
		return
	}

	list, err := service.ListHistory(id)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	server.Ok(list, w)
}

func DeleteHistoryHandler(w http.ResponseWriter, r *http.Request) {
	info := auth(w, r)
	if info == nil {
		return
	}

	historyId, err := server.ReadPathUUID(HistoryIdKey, r)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	history, err := service.GetHistory(historyId)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if history == nil {
		server.NotFound(w)
		return
	}

	if !uuid.Equals(*history.AuthId, info.CredentialsId) {
		log.Printf("access denied %s -> %s", info.CredentialsId, history.AuthId)
		server.Forbidden(w)
		return
	}

	ok, err := service.DeleteFromHistory(historyId)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if ok {
		server.Ok(nil, w)
		return
	}

	server.NotFound(w)
}
