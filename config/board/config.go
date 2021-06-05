package board

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Database *psql.DatabaseConfig
var Http *config.HttpConfig

func init() {
	Service = config.Service("board")
	if Service == nil {
		panic("board config required!")
	}

	Database = Service.Database.PsqlConfig()
	if Database == nil {
		panic("board database config required!")
	}

	Http = Service.HttpConfig
	if Http == nil {
		panic("board http config required!")
	}
}
