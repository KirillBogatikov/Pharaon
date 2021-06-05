package auth

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	"github.com/Projector-Solutions/Pharaon-config/embed"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

type Config struct {
	SecretKey string `yaml:"secret_key"`
}

var Service *config.ServiceConfig
var Database *psql.DatabaseConfig
var Env *Config

func init() {
	Service = config.Service("auth")
	if Service == nil {
		panic("auth config not found!")
	}

	Database = Service.Database.PsqlConfig()

	Env = &Config{}
	err := config.LoadConfig(embed.AuthDev, embed.AuthProd, Env)
	if err != nil {
		panic(err)
	}
}
