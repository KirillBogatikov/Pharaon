package user

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	"github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
)

var (
	ApiURL = ""
)

func init() {
	if ApiURL = config.Service.HttpConfig.ApiUrl; len(ApiURL) == 0 {
		panic("user service url required")
	}
}

type Credentials struct {
	Id    *uuid.UUID `json:"id"`
	Login string     `json:"login"`
	Email string     `json:"email"`
}

type Name struct {
	First      string `json:"first"`
	Last       string `json:"last"`
	Patronymic string `json:"patronymic"`
}

type PersonalData struct {
	Id        *uuid.UUID `json:"id"`
	Phone     string     `json:"phone"`
	BirthDate *time.Date `json:"birth_date"`
	Photo     *uuid.UUID `json:"photo_id,omitempty"`
	Name      *Name      `json:"name"`
}

type User struct {
	Id          *uuid.UUID
	Credentials *Credentials
	Personal    *PersonalData
}
