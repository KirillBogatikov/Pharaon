package middleware

import (
	"net/http"
)

const (
	ContentTypeHeader = "Content-Type"
	ContentTypeJson   = "application/json"
)

func JsonTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		header := w.Header()
		if ct := header.Get(ContentTypeHeader); len(ct) == 0 {
			header.Add(ContentTypeHeader, ContentTypeJson)
		}
		handler.ServeHTTP(w, request)
	})
}
