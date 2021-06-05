package jwt

import (
	ptime "github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	Id            *uuid.UUID      `json:"jti"`
	IssuedAt      *ptime.DateTime `json:"iat"`
	ExpiresAt     *ptime.DateTime `json:"exp"`
	CredentialsId *uuid.UUID      `json:"credentials"`
}

func NewClaims(credentialsId *uuid.UUID) *Claims {
	id := uuid.New()
	now := time.Now().UTC()

	return &Claims{
		Id:            &id,
		IssuedAt:      ptime.NewDateTime(now),
		ExpiresAt:     ptime.NewDateTime(now.Add(tokenLifeTime)),
		CredentialsId: credentialsId,
	}
}

func (c *Claims) Valid() error {
	if c.Id == nil || c.CredentialsId == nil {
		return TokenIncorrectError
	}
	exp := c.ExpiresAt.UnixNano()
	now := time.Now().UTC().UnixNano()
	if exp < now {
		return TokenIncorrectError
	}

	return nil
}
