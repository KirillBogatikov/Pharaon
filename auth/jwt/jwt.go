package jwt

import (
	"errors"
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	tokenLifeTime = time.Hour * 24 * 7
)

var (
	TokenIncorrectError = errors.New("token incorrect")
)

func GenerateToken(credentialsId uuid.UUID) (string, error) {
	claims := NewClaims(&credentialsId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.SecretKey))
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := NewClaims(nil)
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Env.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if tkn.Valid {
		return claims, nil
	}

	return nil, TokenIncorrectError
}
