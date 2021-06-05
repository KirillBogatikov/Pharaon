package files

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Http *config.HttpConfig
var Database *psql.DatabaseConfig

func init() {
	Service = config.Service("files")
	if Service == nil {
		panic("files config not found!")
	}

	Http = Service.HttpConfig
	Database = Service.Database.PsqlConfig()
}
