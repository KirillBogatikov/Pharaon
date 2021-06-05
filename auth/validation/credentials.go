package validation

import (
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"pharaon-auth/data"
)

var v = &validation.Validator{}

func ValidatePassword(password string) validation.FieldResult {
	return v.ValidateString(password, PasswordRule)
}

func ValidateCredentials(credentials *data.Credentials, ignorePassword bool) *ModelResult {
	result := &ModelResult{
		Login:    v.ValidateString(credentials.Login, LoginRule),
		Password: ValidatePassword(credentials.Password),
		Email:    v.ValidateString(credentials.Email, EmailRule),
	}

	if ignorePassword {
		result.Password = validation.Valid
	}

	return result
}
