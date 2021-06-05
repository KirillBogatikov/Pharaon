package service

import (
	"log"
	"net/http"
	"pharaon-user/auth"
	"pharaon-user/errors"
	"pharaon-user/model"
	"pharaon-user/personal"
	"pharaon-user/validation"
)

func Update(user *model.User) (bool, *validation.UserResult, error) {
	result, err := validation.Validate(user)
	if err != nil {
		return false, nil, err
	}

	if result == nil {
		return false, nil, nil
	}

	if !result.IsValid() {
		return false, result, nil
	}

	oldUser, err := repo.GetById(*user.Id)
	if err != nil {
		return false, nil, err
	}
	if oldUser != nil {
		return false, nil, nil
	}

	transaction := OpenTransaction()
	err = transaction.Exec(func() error {
		status, _, err := auth.Update(user.Credentials)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		log.Printf("ERROR: Data inconsistent! Credentials updated and can not be canceled. Check user %s\n", user.Id)
		return nil
	})
	if err != nil {
		return false, nil, err
	}

	err = transaction.Exec(func() error {
		status, _, err := personal.Update(user.Personal)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		log.Printf("ERROR: Data inconsistent! Personal Data updated and can not be canceled. Check user %s\n", user.Id)
		return nil
	})
	if err != nil {
		return false, nil, err
	}

	return true, nil, nil
}
