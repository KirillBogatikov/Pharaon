package service

import "pharaon-migration/data"

var initRepo *data.InitRepository
var serviceRepo *data.ServiceRepository

func InitRepository() (err error) {
	initRepo, err = data.NewInitRepository()
	if err != nil {
		return
	}

	serviceRepo, err = data.NewServiceRepository()
	if err != nil {
		return
	}

	return
}
