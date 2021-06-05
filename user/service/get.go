package service

import (
	"github.com/google/uuid"
	"net/http"
	"pharaon-user/auth"
	"pharaon-user/errors"
	"pharaon-user/model"
	"pharaon-user/personal"
)

func GetById(id uuid.UUID) (*model.User, error) {
	user, err := repo.GetById(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	status, credentials, err := auth.Get(*user.Credentials.Id)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.UnexpectedStatusError(status)
	}
	user.Credentials = credentials

	status, personalData, err := personal.Get(*user.Personal.Id)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.UnexpectedStatusError(status)
	}
	user.Personal = personalData

	return user, nil
}
