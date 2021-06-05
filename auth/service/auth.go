package service

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"pharaon-auth/jwt"
	"strings"
)

func Auth(header string) (*AuthInfo, bool, error) {
	if strings.Index(header, "Bearer ") != 0 {
		return nil, false, nil
	}

	claims, err := jwt.ParseToken(header[7:])
	if err != nil {
		if err.Error() == jwt2.ErrSignatureInvalid.Error() ||
			err.Error() == jwt.TokenIncorrectError.Error() {
			return nil, false, err
		}

		return nil, true, err
	}

	return &AuthInfo{
		CredentialsId: *claims.CredentialsId,
	}, false, nil
}
