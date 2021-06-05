package service

import (
	validation2 "github.com/Projector-Solutions/Pharaon-tools/validation"
	"github.com/google/uuid"
	"pharaon-auth/data"
	"pharaon-auth/encryption"
	"pharaon-auth/jwt"
	"pharaon-auth/validation"
)

var (
	defaultMethod = encryption.BCRYPT
)

func Signup(login, password, email string) (*validation.ModelResult, *data.Credentials, error) {
	auth := &data.Credentials{
		Login:    login,
		Password: password,
		Email:    email,
		Method:   string(defaultMethod),
	}

	val := validation.ValidateCredentials(auth, false)
	if !val.IsValid() {
		return val, nil, nil
	}

	exists, err := authRepo.GetByLogin(login)
	if err != nil {
		return nil, nil, err
	}

	if exists != nil {
		val.Login = validation2.Busy
		return val, nil, nil
	}

	authId := uuid.New()
	hash, err := encryption.GetInstance(defaultMethod).Encrypt(password)
	if err != nil {
		return nil, nil, err
	}

	auth.Id = &authId
	auth.Password = hash

	err = authRepo.Insert(*auth)
	if err != nil {
		return nil, nil, err
	}

	auth.Password = ""
	auth.Method = ""

	return val, auth, nil
}

func Login(ip, login, password string) (*data.Credentials, error) {
	c, err := authRepo.GetByLogin(login)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, NotFoundError
	}

	enc := encryption.GetInstance(encryption.Method(c.Method))
	passwordCheck, err := enc.Compare(c.Password, password)
	if err != nil {
		return nil, err
	}
	if !passwordCheck {
		return nil, PasswordIncorrectError
	}

	token, err := jwt.GenerateToken(*c.Id)
	if err != nil {
		return nil, err
	}

	c.JWT = token

	_, err = SaveToHistory(*c.Id, ip)
	if err != nil {
		return nil, err
	}

	c.Hide()

	return c, nil
}

func GetCredentials(id uuid.UUID) (*data.Credentials, error) {
	c, err := authRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, NotFoundError
	}

	c.Hide()

	return c, nil
}

func Merge(credentials *data.Credentials) (bool, error) {
	origin, err := authRepo.GetById(*credentials.Id)
	if err != nil {
		return false, err
	}

	if origin == nil {
		return false, nil
	}

	if len(credentials.Login) == 0 {
		credentials.Login = origin.Login
	}

	if len(credentials.Password) == 0 {
		credentials.Password = origin.Password
	}

	if len(credentials.Email) == 0 {
		credentials.Email = origin.Email
	}

	credentials.Method = origin.Method
	return true, nil
}

func UpdateCredentials(credentials *data.Credentials) (bool, *validation.ModelResult, error) {
	needHash := len(credentials.Password) > 0

	found, err := Merge(credentials)
	if err != nil {
		return false, nil, err
	}
	if !found {
		return false, nil, nil
	}

	result := validation.ValidateCredentials(credentials, !needHash)
	if !result.IsValid() {
		return false, result, nil
	}

	if needHash {
		originPassword := credentials.Password

		hash, err := encryption.GetInstance(defaultMethod).Encrypt(originPassword)
		if err != nil {
			return false, result, err
		}

		credentials.Password = hash

		defer func() {
			credentials.Password = originPassword
		}()
	}

	found, err = authRepo.Update(*credentials)
	return found, result, err
}

func DeleteAuth(id uuid.UUID) (bool, error) {
	return authRepo.Delete(id)
}
