package user

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Http *config.HttpConfig
var Database *psql.DatabaseConfig

func init() {
	Service = config.Service("user")
	if Service == nil {
		panic("user config not found!")
	}

	Http = Service.HttpConfig
	if Http == nil {
		panic("user http config required!")
	}

	Database = Service.Database.PsqlConfig()
}
