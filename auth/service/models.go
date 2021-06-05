package service

import "github.com/google/uuid"

type AuthInfo struct {
	CredentialsId uuid.UUID `json:"id"`
}
