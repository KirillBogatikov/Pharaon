package sql

import _ "embed"

//go:embed service/getByName.sql
var ServiceGetByName string

//go:embed service/insert.sql
var ServiceInsert string

//go:embed service/update.sql
var ServiceUpdate string

//go:embed service/delete.sql
var ServiceDelete string

//go:embed init/check.sql
var InitCheck string

//go:embed init/create.sql
var InitCreate string

//go:embed init/init_service.sql
var InitService string
