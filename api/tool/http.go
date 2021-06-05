package tool

import (
	"errors"
	"net/http"
	"time"
)

func DefaultHttpClient() *http.Client {
	return &http.Client{}
}

func LongTimeoutHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
	}
}

var (
	UnexpectedStatusError = errors.New("unexpected response status code")
)
