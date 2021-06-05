package header

import (
	"net/http"
	"strings"
)

func getUpperBound(path string) int {
	if i := strings.Index(path[1:], "/"); i > -1 {
		return i + 1
	}

	return len(path)
}

func GetProxyService(r *http.Request) (string, string) {
	index := getUpperBound(r.RequestURI)
	return r.RequestURI[1:index], r.RequestURI[index+1:]
}
