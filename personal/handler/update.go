package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-personal/model"
	"pharaon-personal/service"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	data := &model.Data{}
	err := server.ReadJson(r, data)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	data.Id = &id

	found, result, err := service.Update(data)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if result == nil || !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(data, w)
}
