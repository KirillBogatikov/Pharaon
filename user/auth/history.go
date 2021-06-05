package auth

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/google/uuid"
	"net/http"
)

func GetHistory(id uuid.UUID) (status int, list []History, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/history/user/%s", authApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		err = hc.ReadJSON(&list)
	}

	return
}

func DeleteFromHistory(id uuid.UUID) (status int, err error) {
	hc := client.NewClient(&http.Request{
		Method: "DELETE",
		URL:    client.MustFormatURL("%s/history/%s", authApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	return
}
