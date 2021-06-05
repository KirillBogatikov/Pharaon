package client

import (
	"net/http"
	"net/url"
)

func RequestJson(method, rawUrl string, requestData interface{}, responseMapping ResponseMapping) (*HttpClient, error) {
	requestUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	req := NewClient(&http.Request{
		Method: method,
		URL:    requestUrl,
	})
	if requestData != nil {
		err = req.WriteJSON(requestData)
		if err != nil {
			return nil, err
		}
	}

	_, err = req.Do(&http.Client{})
	if err != nil {
		return req, err
	}

	err = req.ReadComplexJSON(responseMapping)
	if err != nil {
		return req, err
	}

	return req, nil
}
