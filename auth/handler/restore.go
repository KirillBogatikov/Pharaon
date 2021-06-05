package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pharaon-auth/service"
)

func StartRestore(w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	token, err := service.CreateRestoreToken(id)
	if err != nil {
		if err == service.TokenExists {
			server.Ok(nil, w)
			return
		}

		server.InternalServerError(w)
		return
	}

	server.Created(token, w)
}

func ApplyRestore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token, ok := vars["token"]
	if !ok || len(token) == 0 {
		server.Forbidden(w)
		return
	}

	password := r.URL.Query().Get("password")

	ok, result, err := service.ApplyRestoreToken(token, password)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if result != nil && !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	if !ok {
		server.Forbidden(w)
	}

	server.Accepted(nil, w)
}
