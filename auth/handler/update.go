package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-auth/data"
	"pharaon-auth/service"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	credentials := &data.Credentials{}

	err = server.ReadJson(r, credentials)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
	}

	credentials.Id = &id
	found, result, err := service.UpdateCredentials(credentials)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	if !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	server.Ok(credentials, w)
}
