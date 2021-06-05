package personal

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Http *config.HttpConfig
var Database *psql.DatabaseConfig

func init() {
	Service = config.Service("personal")
	if Service == nil {
		panic("personal config not found!")
	}

	Http = Service.HttpConfig
	Database = Service.Database.PsqlConfig()
}
