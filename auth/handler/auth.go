package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	info := auth(w, r)
	if info == nil {
		return
	}

	server.Ok(info, w)
}
