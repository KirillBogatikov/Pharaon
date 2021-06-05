package service

import (
	"pharaon-auth/data"
	"pharaon-auth/validation"
)

func Validate(credentials *data.Credentials) (bool, *validation.ModelResult, error) {
	ignorePassword := len(credentials.Password) > 0

	if credentials.Id != nil {
		found, err := Merge(credentials)
		if err != nil {
			return false, nil, err
		}

		if !found {
			return false, nil, nil
		}
	}

	result := validation.ValidateCredentials(credentials, ignorePassword)
	return true, result, nil
}
