package errors

import "fmt"

func UnexpectedStatusError(status int) error {
	return fmt.Errorf("personal service unexpected response: %d", status)
}
