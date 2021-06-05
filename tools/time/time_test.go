package time

import (
	"encoding/json"
	"testing"
	gotime "time"
)

type DateTimeHolder struct {
	Start *DateTime `json:"start"`
}

type DateHolder struct {
	Date *Date `json:"date"`
}

type TimeHolder struct {
	Time *Time `json:"time"`
}

func nullTime(t *testing.T) {
	th := &DateTimeHolder{}
	th.Start = NewZeroDateTime()

	data, err := json.Marshal(th)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != `{"start":null}` {
		t.Fatal("json incorrect")
	}
}

func dateTest(t *testing.T) {
	th := &DateHolder{}

	date := gotime.Date(2020, 03, 21, 0, 0, 0, 0, gotime.UTC)
	th.Date = NewDate(date)

	data, err := json.Marshal(th)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != `{"date":"2020-03-21"}` {
		t.Fatal("json incorrect", string(data))
	}
}

func timeTest(t *testing.T) {
	th := &TimeHolder{}

	date := gotime.Date(2020, 03, 21, 21, 18, 5, 0, gotime.UTC)
	th.Time = NewTime(date)

	data, err := json.Marshal(th)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != `{"time":"21:18:05"}` {
		t.Fatal("json incorrect", string(data))
	}
}

func array(t *testing.T) {
	origin := `[{"start":"2021-05-10 19:29:36"},{"start":"2021-04-10 19:29:36"},{"start":"2020-05-10 19:29:36"}]`
	var ths []DateTimeHolder

	err := json.Unmarshal([]byte(origin), &ths)
	if err != nil {
		t.Fatal(err)
	}

	date := gotime.Date(2021, 05, 10, 19, 29, 36, 0, gotime.UTC)
	if actual := ths[0].Start.Time; actual != date {
		t.Fatal("expected", date, "received", actual)
	}

	bytes, err := json.Marshal(ths)
	if err != nil {
		t.Fatal(err)
	}

	if origin != string(bytes) {
		t.Fatal("expected", origin, "received", string(bytes))
	}
}

func TestTime(t *testing.T) {
	t.Run("null", nullTime)
	t.Run("date", dateTest)
	t.Run("time", timeTest)
	t.Run("datetime array", array)
}
