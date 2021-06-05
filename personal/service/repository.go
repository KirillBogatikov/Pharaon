package service

import "pharaon-personal/data"

var repo *data.Repository

func InitRepository() (err error) {
    repo, err = data.NewRepository()
    return
}