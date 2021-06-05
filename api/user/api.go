package user

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/google/uuid"
	"net/http"
)

func GetById(id uuid.UUID, token string) (status int, u *User, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/%s", ApiURL, id.String()),
	})
	hc.UseAuthorizationToken(token)

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		u = &User{}
		err = hc.ReadJSON(u)
		return
	}

	return
}

func Check(id []uuid.UUID, token string) (status int, result []bool, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/check", ApiURL),
	})
	hc.UseAuthorizationToken(token)

	err = hc.WriteJSON(id)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		result = make([]bool, 0)
		err = hc.ReadJSON(&result)
	}

	return
}
