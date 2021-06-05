package time

import (
	"database/sql/driver"
	"time"
)

const DateFormat = "yyyy-MM-dd"

var DateLayout = ParseFormat(DateFormat)

type Date struct {
	time.Time
}

func NewDate(t time.Time) *Date {
	return &Date{t}
}

func NewZeroDate() *Date {
	return &Date{}
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = unmarshalJSON(DateLayout, data)
	return
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return marshalJSON(DateLayout, d.Time), nil
}

func (d *Date) Value() (driver.Value, error) {
	return sqlValue(d.Time), nil
}

func (d *Date) Scan(value interface{}) (err error) {
	d.Time, err = sqlScan(value)
	return
}
