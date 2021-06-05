package handler

import (
	tool "github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-migration/data"
	"pharaon-migration/service"
)

func RegisterServiceHandler(w http.ResponseWriter, r *http.Request) {
	d := &data.ServiceData{}
	err := tool.ReadJson(r, d)
	if err != nil {
		tool.BadRequest(nil, w)
		return
	}

	serviceInfo, err := service.Register(d)
	if err == data.ServiceNotFound {
		tool.NotFound(w)
		return
	}

	if err != nil {
		log.Println(err)
		tool.InternalServerError(w)
		return
	}

	tool.Ok(serviceInfo, w)
}
