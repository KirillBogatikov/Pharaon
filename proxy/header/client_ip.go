package header

import (
	"net"
	"net/http"
)

func GetClientAddr(r *http.Request) string {
	header := r.Header
	if header != nil {
		if xForwardedIp := getXForwardedFor(header); len(xForwardedIp) > 0 {
			return xForwardedIp
		}

		if xProxyUserId := header.Get(XProxyUserId); len(xProxyUserId) > 0 {
			return xProxyUserId
		}
	}

	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

func getXForwardedFor(header http.Header) string {
	xInfo := header.Values(XForwardedFor)
	if xInfo == nil || len(xInfo) == 0 {
		return ""
	}

	return xInfo[0]
}
