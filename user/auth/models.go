package auth

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	ptime "github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"github.com/google/uuid"
)

var (
	authApiURL = ""
)

func init() {
	if authApiURL = config.Service.HttpConfig.ApiUrl; len(authApiURL) == 0 {
		panic("auth API URL required")
	}
}

type Credentials struct {
	Id       *uuid.UUID `json:"id"`
	Login    string     `json:"login"`
	Password string     `json:"password"`
	Email    string     `json:"email"`
	JWT      string     `json:"jwt"`
}

type History struct {
	Id   *uuid.UUID      `json:"id"`
	Ip   string          `json:"ip"`
	Time *ptime.DateTime `json:"time"`
}

type RestoreToken struct {
	Id      *uuid.UUID      `json:"id"`
	Token   string          `json:"token"`
	Expires *ptime.DateTime `json:"expires"`
}

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
