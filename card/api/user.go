package api

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-api/user"
	"github.com/google/uuid"
	"net/http"
)

func GetUserById(userId uuid.UUID, token string) (*user.User, error) {
	status, u, err := user.GetById(userId, token)
	if err != nil {
		return nil, err
	}

	switch status {
	case http.StatusNotFound:
		return nil, nil
	case http.StatusOK:
		return u, nil
	}

	return nil, tool.UnexpectedStatusError
}
