package client

import (
	"fmt"
	"net/url"
)

func FormatURL(format string, args ...interface{}) (*url.URL, error) {
	return url.Parse(fmt.Sprintf(format, args...))
}

func MustFormatURL(format string, args ...interface{}) *url.URL {
	u, err := FormatURL(format, args...)
	if err != nil {
		panic(err)
	}

	return u
}
