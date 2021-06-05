package auth

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"net/http"
)

func Validate(c *Credentials) (status int, m *ModelResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/validate", authApiURL),
	})

	err = hc.WriteJSON(c)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		m = &ModelResult{}
		err = hc.ReadJSON(m)
	}

	return
}
