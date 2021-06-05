package auth

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/google/uuid"
	"net/http"
)

func Get(id uuid.UUID) (status int, c *Credentials, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/user/%s", authApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		c := &Credentials{}
		err = hc.ReadJSON(c)
	}

	return
}

func Update(c *Credentials) (status int, m *ModelResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/user/%s", authApiURL, c.Id),
	})

	err = hc.WriteJSON(c)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusBadRequest {
		m = &ModelResult{}
		err = hc.ReadJSON(m)
	} else if status == http.StatusOK {
		err = hc.ReadJSON(c)
	}

	return
}

func Delete(id uuid.UUID) (status int, err error) {
	hc := client.NewClient(&http.Request{
		Method: "DELETE",
		URL:    client.MustFormatURL("%s/user/%s", authApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	return
}
