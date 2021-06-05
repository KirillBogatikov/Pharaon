package validation

import (
	"net/http"
	"pharaon-user/auth"
	"pharaon-user/errors"
	"pharaon-user/model"
	"pharaon-user/personal"
)

func ValidateCredentials(credentials *auth.Credentials) (*auth.ModelResult, error) {
	status, cResult, err := auth.Validate(credentials)
	if err != nil {
		return nil, err
	}

	if status == http.StatusNotFound {
		return nil, nil
	}

	if status != http.StatusOK {
		return nil, errors.UnexpectedStatusError(status)
	}

	return cResult, nil
}

func ValidateData(data *personal.Data) (*personal.DataResult, error) {
	status, pResult, err := personal.Validate(data)
	if err != nil {
		return nil, err
	}

	if status == http.StatusNotFound {
		return nil, nil
	}

	if status != http.StatusOK {
		return nil, errors.UnexpectedStatusError(status)
	}

	return pResult, nil
}

func Validate(user *model.User) (*UserResult, error) {
	cResult, err := ValidateCredentials(user.Credentials)
	if err != nil {
		return nil, err
	}

	if cResult == nil {
		return nil, nil
	}

	pResult, err := ValidateData(user.Personal)
	if err != nil {
		return nil, err
	}

	if pResult == nil {
		return nil, nil
	}

	return &UserResult{
		Credentials: cResult,
		Personal:    pResult,
	}, nil
}
