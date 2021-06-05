package service

import (
	"crypto/rand"
	"errors"
	ptime "github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
	"github.com/robfig/cron"
	"log"
	"math/big"
	"pharaon-auth/data"
	"pharaon-auth/validation"
	"time"
)

var (
	tokenAlphabet       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	tokenAlphabetLength = big.NewInt(int64(len(tokenAlphabet)))
	TokenExists         = errors.New("password restoring already requested")
)

const (
	RestoreTokenLength = 64
)

func RandomToken(length int) (string, error) {
	token := make([]rune, length)
	for i := range token {
		num, err := rand.Int(rand.Reader, tokenAlphabetLength)
		if err != nil {
			return "", err
		}

		token[i] = tokenAlphabet[num.Int64()]
	}

	return string(token), nil
}

func CreateRestoreToken(id uuid.UUID) (*data.RestoreToken, error) {
	existToken, err := tokenRepo.GetByAuth(id)
	if err != nil {
		return nil, err
	}

	if existToken != nil {
		existToken.Token = ""
		return existToken, TokenExists
	}

	tokenId := uuid.New()
	tokenText, err := RandomToken(RestoreTokenLength)
	if err != nil {
		return nil, err
	}

	token := data.RestoreToken{
		Id:      &tokenId,
		AuthId:  &id,
		Token:   tokenText,
		Expires: ptime.NewDateTime(time.Now().Add(time.Hour * 24)),
	}

	err = tokenRepo.Insert(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func ApplyRestoreToken(tokenText string, password string) (bool, *validation.ModelResult, error) {
	token, err := tokenRepo.GetByToken(tokenText)
	if err != nil {
		return false, nil, err
	}

	_, err = tokenRepo.Delete(*token.Id)
	if err != nil {
		return false, nil, nil
	}

	if token.Expires.UnixNano() < time.Now().UnixNano() {
		return false, nil, err
	}

	auth := &data.Credentials{
		Id:       token.AuthId,
		Password: password,
	}

	found, result, err := UpdateCredentials(auth)
	return found, result, err
}

func StartAutoClearing() error {
	c := cron.New()
	err := c.AddFunc("0 0 */12 * * *", func() {
		log.Println("Start clearing old tokens...")
		count, err := tokenRepo.Clear()
		if err != nil {
			log.Println("Clearing failed with ", err)
		}

		log.Printf("Deleted %d tokens\n", count)
	})

	if err != nil {
		return err
	}

	go c.Start()
	return nil
}
