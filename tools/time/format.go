package time

import (
	"database/sql/driver"
	"strings"
	"time"
)

func replace(origin string, params ...string) string {
	for i := 0; i < len(params); i += 2 {
		origin = strings.ReplaceAll(origin, params[i], params[i+1])
	}

	return origin
}

func ParseFormat(format string) string {
	return replace(format, "yyyy", "2006", "yy", "06",
		"MMMM", "January", "MMM", "Jan", "MM", "01", "M", "1",
		"dd", "02", "d", "2",
		"EEE", "Mon", "EEEEEE", "Monday",
		"HH", "15", "hh", "03", "hh", "03", "h", "3",
		"mm", "04", "m", "4",
		"ss", "05", "s", "5",
		"z", "MST")
}

type FormattedTime struct {
	Format string
	time.Time
}

func NewFormattedTime(format string, t time.Time) *FormattedTime {
	return &FormattedTime{format, t}
}

func NewZeroFormattedTime(format string) *FormattedTime {
	return &FormattedTime{Format: format}
}

func (t *FormattedTime) Layout() string {
	return ParseFormat(t.Format)
}

func (t *FormattedTime) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = unmarshalJSON(t.Layout(), data)
	return
}

func (t *FormattedTime) MarshalJSON() ([]byte, error) {
	return marshalJSON(t.Layout(), t.Time), nil
}

func (t *FormattedTime) Value() (driver.Value, error) {
	return sqlValue(t.Time), nil
}

func (t *FormattedTime) Scan(value interface{}) (err error) {
	t.Time, err = sqlScan(value)
	return
}
