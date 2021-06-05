package service

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"pharaon-user/auth"
	"pharaon-user/errors"
	"pharaon-user/personal"
)

func Delete(id uuid.UUID) (bool, error) {
	user, err := repo.GetById(id)
	if err != nil {
		return false, err
	}

	transaction := OpenTransaction()

	err = transaction.Exec(func() error {
		status, err := auth.Delete(*user.Credentials.Id)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		log.Printf("ERROR: Data inconsistent! Credentials deleted and can not be canceled. Check user %s\n", user.Id)
		return nil
	})
	if err != nil {
		return false, err
	}

	err = transaction.Exec(func() error {
		status, err := personal.Delete(*user.Personal.Id)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return errors.UnexpectedStatusError(status)
		}

		return nil
	}, func() error {
		log.Printf("ERROR: Data inconsistent! Personal Data deleted and can not be canceled. Check user %s\n", user.Id)
		return nil
	})
	if err != nil {
		return false, err
	}

	found, err := repo.Delete(*user.Id)
	return found, err
}
