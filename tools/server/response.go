package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {
	WriteJson(nil, http.StatusInternalServerError, w)
}

func BadRequest(data interface{}, w http.ResponseWriter) {
	WriteJson(data, http.StatusBadRequest, w)
}

func NotFound(w http.ResponseWriter) {
	WriteJson(nil, http.StatusNotFound, w)
}

func Unauthorized(w http.ResponseWriter) {
	WriteJson(nil, http.StatusUnauthorized, w)
}

func Forbidden(w http.ResponseWriter) {
	WriteJson(nil, http.StatusForbidden, w)
}

func Ok(data interface{}, w http.ResponseWriter) {
	WriteJson(data, http.StatusOK, w)
}

func Created(data interface{}, w http.ResponseWriter) {
	WriteJson(data, http.StatusCreated, w)
}

func Accepted(data interface{}, w http.ResponseWriter) {
	WriteJson(data, http.StatusAccepted, w)
}

func WriteJson(data interface{}, status int, w http.ResponseWriter) {
	if data == nil {
		w.WriteHeader(status)
		return
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		InternalServerError(w)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
