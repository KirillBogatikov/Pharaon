package sql

import _ "embed"

//go:embed card/get_by_id.sql
var CardGetById string

//go:embed card/get_in_ids.sql
var CardGetInIds string

//go:embed card/insert.sql
var CardInsert string

//go:embed card/update.sql
var CardUpdate string

//go:embed card/delete.sql
var CardDelete string
