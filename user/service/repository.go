package service

import "pharaon-user/data"

var repo *data.Repository

func InitRepository() (err error) {
	repo, err = data.NewRepository()
	return err
}
