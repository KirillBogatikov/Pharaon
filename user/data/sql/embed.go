package sql

import _ "embed"

//go:embed user/check_exists.sql
var UserCheckExists string

//go:embed user/get_by_id.sql
var UserGetById string

//go:embed user/insert.sql
var UserInsert string

//go:embed user/delete.sql
var UserDelete string
