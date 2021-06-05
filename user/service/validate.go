package service

import (
	"pharaon-user/model"
	"pharaon-user/validation"
)

func Validate(user *model.User) (bool, *validation.UserResult, error) {
	result, err := validation.Validate(user)
	if err != nil {
		return false, nil, err
	}

	if result == nil {
		return false, nil, nil
	}

	return true, result, nil
}
