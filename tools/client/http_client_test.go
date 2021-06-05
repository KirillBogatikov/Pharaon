package client

import (
	"fmt"
	"testing"
)

func post(t *testing.T) {
	pet := &Pet{}
	req, err := RequestJson("POST", apiPostUrl, natasha, ResponseMapping{
		200: pet,
	})
	if err != nil {
		t.Fatal(err)
	}

	if code := req.Response().StatusCode; code != 200 {
		t.Fatal("expected status 200, actual is ", code, string(req.ResponseBody()))
	}

	natasha.Id = pet.Id
	if !equals(natasha, pet) {
		t.Fatal("expected ", natasha, " received ", pet)
	}
}

func get(t *testing.T) {
	req, err := RequestJson("GET", fmt.Sprintf(apiUrl, natasha.Id), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if code := req.Response().StatusCode; code != 200 {
		t.Fatal("expected status 200, actual is ", code, string(req.ResponseBody()))
	}

	pet := &Pet{}
	err = req.ReadJSON(pet)
	if err != nil {
		t.Fatal(err)
	}

	if !equals(natasha, pet) {
		t.Fatal("expected ", natasha, " received ", pet)
	}
}

func deleteNatasha(t *testing.T) {
	req, err := RequestJson("DELETE", fmt.Sprintf(apiUrl, natasha.Id), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if code := req.Response().StatusCode; code != 200 {
		t.Fatal("expected status 200, actual is ", code, string(req.ResponseBody()))
	}
}

func notFound(t *testing.T) {
	req, err := RequestJson("GET", fmt.Sprintf(apiUrl, natasha.Id), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if code := req.Response().StatusCode; code != 404 {
		t.Fatal("expected status 404, actual is ", code, string(req.ResponseBody()))
	}
}

func TestRequestJson(t *testing.T) {
	t.Run("Post Natasha", post)
	t.Run("Get Natasha", get)
	t.Run("Delete Natasha", deleteNatasha)
	t.Run("Natasha not found", notFound)
	t.Run("Complex get", complexGet)
}
