package service

import "errors"

var (
	PasswordIncorrectError = errors.New("password incorrect")
	NotFoundError          = errors.New("not found")
)
