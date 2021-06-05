package time

import (
	"fmt"
	"strings"
	"time"
)

func unmarshalJSON(layout string, data []byte) (t time.Time, err error) {
	word := strings.Trim(string(data), "\"")
	if word == "null" {
		return
	}

	t, err = time.Parse(layout, word)
	return
}

func marshalJSON(layout string, t time.Time) []byte {
	if t.IsZero() {
		return []byte("null")
	}

	word := fmt.Sprintf(`"%s"`, t.Format(layout))
	return []byte(word)
}
