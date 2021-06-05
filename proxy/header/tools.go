package header

import (
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"net/http"
	"strconv"
	"strings"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XProxyUserId  = "X-ProxyUser-Ip"
)

var hideHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func CopyHeader(from, into http.Header) {
	for name, values := range from {
		for _, v := range values {
			into.Add(name, v)
		}
	}
}

func ClearHeader(header http.Header) {
	for _, h := range hideHeaders {
		header.Del(h)
	}
}

func AppendProxyHost(header http.Header, host string) {
	if prior := header.Values(XForwardedFor); prior != nil && len(prior) > 0 {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set(XForwardedFor, host)
}

func GetContentLength(header http.Header) int {
	if header := header.Get(client.ContentTypeHeader); len(header) > 0 {
		i, _ := strconv.Atoi(header)
		return i
	}

	return 0
}

func GetHeadersMap(header http.Header) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range header {
		result[k] = strings.Join(v, ", ")
	}

	return result
}
