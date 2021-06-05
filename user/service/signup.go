package service

import (
	"net/http"
	"pharaon-user/auth"
	"pharaon-user/errors"
	"pharaon-user/model"
	"pharaon-user/personal"
	"pharaon-user/validation"
)

func Signup(user *model.User) (*validation.UserResult, error) {
	result, err := validation.Validate(user)
	if err != nil {
		return nil, err
	}

	if !result.IsValid() {
		return result, nil
	}

	transaction := OpenTransaction()

	err = transaction.Exec(func() error {
		status, _, err := auth.Signup(user.Credentials)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		status, err := auth.Delete(*user.Credentials.Id)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = transaction.Exec(func() error {
		status, _, err := personal.Signup(user.Personal)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		status, err := personal.Delete(*user.Personal.Id)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = transaction.Exec(func() error {
		return repo.Insert(user)
	}, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
