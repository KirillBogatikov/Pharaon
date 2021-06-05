package card

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Database *psql.DatabaseConfig
var Http *config.HttpConfig

func init() {
	Service = config.Service("card")
	if Service == nil {
		panic("card config required!")
	}

	Database = Service.Database.PsqlConfig()
	if Database == nil {
		panic("card database config required!")
	}

	Http = Service.HttpConfig
	if Http == nil {
		panic("card http config required!")
	}
}
