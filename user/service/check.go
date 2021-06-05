package service

import "github.com/google/uuid"

func CheckExists(ids []uuid.UUID) ([]bool, error) {
	result, err := repo.Check(ids...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
