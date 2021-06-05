package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/Projector-Solutions/Pharaon-tools/uuid"
	"net/http"
	"pharaon-user/model"
	"pharaon-user/service"
)

func UpdateHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(UserIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	if !uuid.Equals(id, i.CredentialsId) {
		//todo: UAC
	}

	user := &model.User{}
	err = server.ReadJson(r, user)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	user.Id = &id
	found, result, err := service.Update(user)
	if err != nil {
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

	server.Ok(user, w)
}
