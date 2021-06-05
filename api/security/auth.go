package security

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"net/http"
)

func Auth(r *http.Request) (status int, info *Info, err error) {
	header := r.Header.Get("Authorization")

	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/auth", ApiURL),
	})

	hc.Request().Header.Set("Authorization", header)

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		info = &Info{}
		err = hc.ReadJSON(info)
	}

	return
}
