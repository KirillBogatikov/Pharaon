package time

import (
	"database/sql/driver"
	"time"
)

const DateTimeFormat = DateFormat + " " + TimeFormat

var DateTimeLayout = ParseFormat(DateTimeFormat)

type DateTime struct {
	time.Time
}

func NewDateTime(t time.Time) *DateTime {
	return &DateTime{t}
}

func NewZeroDateTime() *DateTime {
	return &DateTime{}
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = unmarshalJSON(DateTimeLayout, data)
	return
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return marshalJSON(DateTimeLayout, d.Time), nil
}

func (d *DateTime) Value() (driver.Value, error) {
	return sqlValue(d.Time), nil
}

func (d *DateTime) Scan(value interface{}) (err error) {
	d.Time, err = sqlScan(value)
	return
}
