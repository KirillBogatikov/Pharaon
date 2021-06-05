package service

import (
	"pharaon-auth/data"
)

var authRepo *data.CredentialsRepository
var tokenRepo *data.TokenRepository
var historyRepo *data.HistoryRepository

func InitRepository() (err error) {
	authRepo, err = data.NewCredentialsRepository()
	if err != nil {
		return
	}

	tokenRepo, err = data.NewTokenRepository()
	if err != nil {
		return
	}

	historyRepo, err = data.NewHistoryRepository()
	if err != nil {
		return
	}

	return
}
