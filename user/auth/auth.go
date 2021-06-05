package auth

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"net/http"
)

func Login(login, password, ip string) (status int, c *Credentials, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/login", authApiURL),
	})
	hc.Request().Header.Set("Pharaon-Client-IP", ip)

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		c = &Credentials{}
		err = hc.ReadJSON(c)
		if err != nil {
			return
		}

		return
	}

	return
}

func Signup(c *Credentials) (status int, m *ModelResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/signup", authApiURL),
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
