package auth

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/google/uuid"
	"net/http"
)

func StartRestore(id uuid.UUID) (status int, t *RestoreToken, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/token/user/%s", authApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusCreated {
		t = &RestoreToken{}
		err = hc.ReadJSON(t)
	}

	return
}

func ApplyRestore(token, password string) (status int, m *ModelResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/token/%s/restore", authApiURL, token),
	})
	hc.SetQuery("password", password)

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusBadRequest {
		m := &ModelResult{}
		err = hc.ReadJSON(m)
	}

	return
}
