package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-auth/service"
)

const (
	UserIdKey    = "user_id"
	HistoryIdKey = "history_id"
)

func auth(w http.ResponseWriter, r *http.Request) *service.AuthInfo {
	header := r.Header.Get("Authorization")

	info, internalError, err := service.Auth(header)
	if internalError {
		log.Println(err)
		server.InternalServerError(w)
		return nil
	}

	if err != nil {
		server.Unauthorized(w)
		return nil
	}

	return info
}
