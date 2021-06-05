package model

import (
	"github.com/google/uuid"
	"pharaon-user/auth"
	"pharaon-user/personal"
)

type User struct {
	Id          *uuid.UUID
	Credentials *auth.Credentials
	Personal    *personal.Data
}
