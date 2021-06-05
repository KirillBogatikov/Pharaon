package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-auth/service"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	ip := r.Header.Get("Pharaon-Client-IP")
	login := query.Get("login")
	if len(login) == 0 {
		server.BadRequest(nil, w)
		return
	}

	password := query.Get("password")
	if len(password) == 0 {
		server.BadRequest(nil, w)
		return
	}

	auth, err := service.Login(ip, login, password)
	switch err {
	case service.NotFoundError:
		server.NotFound(w)
		return
	case service.PasswordIncorrectError:
		server.Forbidden(w)
		return
	}
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	server.Ok(auth, w)
}
