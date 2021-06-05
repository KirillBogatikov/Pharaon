package security

import (
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
)

type SecureHandlerFunc func(info Info, w http.ResponseWriter, r *http.Request)

type SecureHandler struct {
	fun SecureHandlerFunc
}

func (s SecureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, info, err := Auth(r)
	if err != nil || status == http.StatusInternalServerError {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	if info == nil {
		server.Unauthorized(w)
		return
	}

	info.Token = r.Header.Get("Authorization")
	s.fun(*info, w, r)
}

func NewSecureHandler(f SecureHandlerFunc) SecureHandler {
	return SecureHandler{fun: f}
}
