package security

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	"github.com/google/uuid"
)

var (
	ApiURL = ""
)

func init() {
	if ApiURL = config.Service.HttpConfig.ApiUrl; len(ApiURL) == 0 {
		panic("auth API URL required")
	}
}

type Info struct {
	CredentialsId uuid.UUID `json:"id"`
	Token         string    `json:"-"`
}
