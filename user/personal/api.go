package personal

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/google/uuid"
	"net/http"
)

func Get(id uuid.UUID) (status int, d *Data, err error) {
	hc := client.NewClient(&http.Request{
		Method: "GET",
		URL:    client.MustFormatURL("%s/user/%s", personalApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		d := &Data{}
		err = hc.ReadJSON(d)
	}

	return
}

func Signup(data *Data) (status int, r *DataResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/user", personalApiURL),
	})

	err = hc.WriteJSON(data)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		d := &Data{}
		err = hc.ReadJSON(d)
	} else if status == http.StatusBadRequest {
		r := &DataResult{}
		err = hc.ReadJSON(r)
	}

	return
}

func Update(data *Data) (status int, r *DataResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "PUT",
		URL:    client.MustFormatURL("%s/user/%s", personalApiURL, data.Id.String()),
	})

	err = hc.WriteJSON(data)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		d := &Data{}
		err = hc.ReadJSON(d)
	} else if status == http.StatusBadRequest {
		r := &DataResult{}
		err = hc.ReadJSON(r)
	}

	return
}

func Delete(id uuid.UUID) (status int, err error) {
	hc := client.NewClient(&http.Request{
		Method: "DELETE",
		URL:    client.MustFormatURL("%s/history/%s", personalApiURL, id.String()),
	})

	status, err = hc.Do(tool.DefaultHttpClient())
	return
}

func Validate(d *Data) (status int, r *DataResult, err error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/validate", personalApiURL),
	})

	err = hc.WriteJSON(d)
	if err != nil {
		return
	}

	status, err = hc.Do(tool.DefaultHttpClient())
	if err != nil {
		return
	}

	if status == http.StatusOK {
		r = &DataResult{}
		err = hc.ReadJSON(r)
	}

	return
}
