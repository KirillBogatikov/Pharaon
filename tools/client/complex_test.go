package client

import (
	"fmt"
	"testing"
)

type PetError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func complexGet(t *testing.T) {
	pet := Pet{}
	petError := PetError{}

	hc, err := RequestJson("GET", fmt.Sprintf(apiUrl, natasha.Id), nil, ResponseMapping{
		200: &pet,
		404: &petError,
	})

	if err != nil {
		t.Fatal(err)
	}

	if code := hc.Response().StatusCode; code != 404 {
		t.Fatal("expected 404, but received ", code, string(hc.ResponseBody()))
	}

	if petError.Code != 1 || petError.Type != "error" || petError.Message != "Pet not found" {
		t.Fatal("expected error")
	}
}
