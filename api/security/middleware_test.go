package security

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"log"
	"net/http"
	"testing"
)

type mockAuthHandler struct {
	t *testing.T
}

func (m *mockAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("Authorization"); token == "hello-world-bearer-token" {
		log.Println(token)

		w.Header().Set(client.ContentTypeHeader, client.ContentTypeJson)
		server.Ok(&Info{}, w)
		return
	}

	server.Unauthorized(w)
}

func testHandler(info Info, w http.ResponseWriter, r *http.Request) {
	log.Println("I'm alive!")
	server.Created(nil, w)
}

func Test(t *testing.T) {
	go func() {
		handler := &mockAuthHandler{t}
		if err := http.ListenAndServe(config.Service.HttpConfig.BindAddress, handler); err != nil {
			t.Fatal(err)
		}
	}()

	go func() {
		if err := http.ListenAndServe(":5000", NewSecureHandler(testHandler)); err != nil {
			t.Fatal(err)
		}
	}()

	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("http://localhost:5000"),
	})

	hc.Request().Header.Set("Authorization", "hello-world-bearer-token")

	status, err := hc.Do(tool.DefaultHttpClient())
	if err != nil {
		t.Fatal(err)
	}

	if status != http.StatusCreated {
		t.Fatal("expected 201, received", status)
	}
}
