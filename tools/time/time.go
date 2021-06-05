package time

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "HH:mm:ss"

var TimeLayout = ParseFormat(TimeFormat)

type Time struct {
	time.Time
}

func NewTime(t time.Time) *Time {
	return &Time{t}
}

func NewZeroTime() *Time {
	return &Time{}
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = unmarshalJSON(TimeLayout, data)
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return marshalJSON(TimeLayout, t.Time), nil
}

func (t *Time) Value() (driver.Value, error) {
	return sqlValue(t.Time), nil
}

func (t *Time) Scan(value interface{}) (err error) {
	t.Time, err = sqlScan(value)
	return
}
