package sql

import _ "embed"

//go:embed data/check_phone.sql
var DataCheckPhone string

//go:embed data/get_by_id.sql
var DataGetById string

//go:embed data/insert.sql
var DataInsert string

//go:embed data/update.sql
var DataUpdate string

//go:embed data/delete.sql
var DataDelete string