package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ContentTypeHeader = "Content-Type"
	ContentTypeJson   = "application/json"
)

var (
	NotJsonError   = errors.New("response content type is not application/json")
	NoMappingError = errors.New("mapping for response status code not found")
)

type ResponseMapping map[int]interface{}

type HttpClient struct {
	response     *http.Response
	request      *http.Request
	responseBody []byte
}

func NewClient(r *http.Request) *HttpClient {
	if r == nil {
		r = &http.Request{}
	}

	if r.Header == nil {
		r.Header = make(http.Header)
	}

	return &HttpClient{request: r}
}

func (h *HttpClient) Do(client *http.Client) (status int, err error) {
	h.response, err = client.Do(h.request)
	if err != nil {
		return
	}

	status = h.response.StatusCode

	h.responseBody, err = ioutil.ReadAll(h.response.Body)
	return
}

func (h *HttpClient) URL() *url.URL {
	return h.request.URL
}

func (h *HttpClient) Query() url.Values {
	return h.request.URL.Query()
}

func (h *HttpClient) GetQuery(key string) string {
	return h.Query().Get(key)
}

func (h *HttpClient) SetQuery(key, value string) {
	h.Query().Set(key, value)
}

func (h *HttpClient) Request() *http.Request {
	return h.request
}

func (h *HttpClient) Response() *http.Response {
	return h.response
}

func (h *HttpClient) ResponseBody() []byte {
	return h.responseBody
}

func (h *HttpClient) ResponseStatus() (int, string) {
	return h.response.StatusCode, h.response.Status
}

func (h *HttpClient) ReadJSON(data interface{}) error {
	if ct := h.response.Header.Get(ContentTypeHeader); ct != ContentTypeJson {
		return NotJsonError
	}

	return json.Unmarshal(h.responseBody, data)
}

func (h *HttpClient) ReadComplexJSON(mapping ResponseMapping) error {
	if mapping == nil {
		return nil
	}

	pointer, ok := mapping[h.response.StatusCode]
	if ok {
		return h.ReadJSON(pointer)
	}

	return NoMappingError
}

func (h *HttpClient) WriteJSON(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	h.request.Header.Set(ContentTypeHeader, ContentTypeJson)
	h.request.Body = ioutil.NopCloser(bytes.NewReader(body))

	return nil
}

func (h *HttpClient) UseAuthorizationToken(token string) {
	h.Request().Header.Set("Authorization", token)
}

func (h *HttpClient) CopyAuthorization(r *http.Request) {
	h.UseAuthorizationToken(r.Header.Get("Authorization"))
}
