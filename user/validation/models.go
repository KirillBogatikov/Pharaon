package validation

import (
	"pharaon-user/auth"
	"pharaon-user/personal"
)

type UserResult struct {
	Credentials *auth.ModelResult
	Personal    *personal.DataResult
}

func (u *UserResult) IsValid() bool {
	return u.Credentials != nil && u.Credentials.IsValid() &&
		u.Personal != nil && u.Personal.IsValid()
}
