package handler

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-auth/data"
	"pharaon-auth/service"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	credentials := &data.Credentials{}

	err := server.ReadJson(r, &credentials)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
	}

	validated, auth, err := service.Signup(credentials.Login, credentials.Password, credentials.Email)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !validated.IsValid() {
		server.BadRequest(validated, w)
		return
	}

	server.Created(auth, w)
}
