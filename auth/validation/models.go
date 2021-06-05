package validation

import "github.com/Projector-Solutions/Pharaon-tools/validation"

type ModelResult struct {
	Login    validation.FieldResult `json:"login"`
	Password validation.FieldResult `json:"password"`
	Email    validation.FieldResult `json:"email"`
}

func (m *ModelResult) IsValid() bool {
	return m.Login == validation.Valid &&
		m.Password == validation.Valid &&
		m.Email == validation.Valid
}

var LoginRule = validation.Rule{
	Min:   4,
	Max:   32,
	Regex: "^[a-zA-Z0-9_!@()~]+$",
}

var PasswordRule = validation.Rule{
	Min:   8,
	Max:   128,
	Regex: "^.+$",
}

var EmailRule = validation.Rule{
	Min:   3,
	Max:   512,
	Regex: `^(\w|[_.$]|-)+@(\w|[_.$]|-)+\.(\w|[_.$]|-)+$`,
}
