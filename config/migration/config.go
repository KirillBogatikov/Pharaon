package migration

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
)

var Service *config.ServiceConfig
var Database *psql.DatabaseConfig

func init() {
	Service = config.Service("migration")
	if Service == nil {
		panic("migration config not found!")
	}

	Database = Service.Database.PsqlConfig()
}
