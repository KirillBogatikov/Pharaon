package handler

import (
	"github.com/Projector-Solutions/Pharaon-api/security"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"pharaon-card/model"
	"pharaon-card/service"
)

func TagListHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	tags, err := service.ListTags()
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	server.Ok(tags, w)
}

func TagAutocompleteHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	query := r.URL.Query().Get("query")
	tags, err := service.SearchTag(query)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	server.Ok(tags, w)
}

func TagCreateHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	//todo: UAC

	tag := &model.Tag{}
	err := server.ReadJson(r, tag)
	if err != nil {
		log.Println(err)
		server.BadRequest(nil, w)
		return
	}

	result, err := service.InsertTag(tag)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !result.IsValid() {
		server.BadRequest(result, w)
		return
	}

	server.Ok(tag, w)
}

func TagDeleteHandler(i security.Info, w http.ResponseWriter, r *http.Request) {
	id, err := server.ReadPathUUID(TagIdKey, r)
	if err != nil {
		server.BadRequest(nil, w)
		return
	}

	found, err := service.DeleteTag(id)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if !found {
		server.NotFound(w)
		return
	}

	server.Ok(nil, w)
}
