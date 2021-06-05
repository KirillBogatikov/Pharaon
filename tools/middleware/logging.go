package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseWriter struct {
	statusCode    int
	contentLength int64
	origin        http.ResponseWriter
}

func NewResponseWriter(origin http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{200, 0, origin}
}

func (r *ResponseWriter) Header() http.Header {
	return r.origin.Header()
}

func (r *ResponseWriter) Write(bytes []byte) (int, error) {
	r.contentLength += int64(len(bytes))
	return r.origin.Write(bytes)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.origin.WriteHeader(statusCode)
}

func (r *ResponseWriter) StatusCode() int {
	return r.statusCode
}

func (r *ResponseWriter) ContentLength() int64 {
	return r.contentLength
}

func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		start := time.Now()

		writer := NewResponseWriter(w)
		handler.ServeHTTP(writer, request)

		ms := time.Now().Sub(start).Milliseconds()
		fmt.Printf("%s %s\n\tStatus: %d\n\tResponse: %d byte\n\tTotal: %dms\n",
			request.Method, request.URL, writer.StatusCode(), writer.ContentLength(), ms)
	})
}
