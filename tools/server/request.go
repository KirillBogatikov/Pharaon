package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadJson(r *http.Request, data interface{}) error {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		return fmt.Errorf("request's Content-Type is %s, application/json required", contentType)
	}

	err = json.Unmarshal(bytes, data)
	return nil
}
