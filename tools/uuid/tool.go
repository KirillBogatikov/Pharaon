package uuid

import "github.com/google/uuid"

func Equals(ids ...uuid.UUID) bool {
	if len(ids) == 0 {
		return true
	}

	firstId := ids[0].String()

	for i := 1; i < len(ids); i++ {
		if ids[i].String() != firstId {
			return false
		}
	}

	return true
}
