package time

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func sqlScan(value interface{}) (t time.Time, err error) {
	switch value.(type) {
	case time.Time:
		t = value.(time.Time)
		break
	case nil:
		break
	default:
		err = fmt.Errorf("unknown time format %v", value)
	}

	return
}

func sqlValue(t time.Time) driver.Value {
	return t
}
