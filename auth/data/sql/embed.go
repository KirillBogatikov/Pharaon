package sql

import _ "embed"

//go:embed history/get.sql
var HistoryGet string

//go:embed history/list.sql
var HistoryList string

//go:embed history/insert.sql
var HistoryInsert string

//go:embed history/delete.sql
var HistoryDelete string

//go:embed restore_token/get_by_auth.sql
var TokenGetByAuth string

//go:embed restore_token/get_by_token.sql
var TokenGetByToken string

//go:embed restore_token/insert.sql
var TokenInsert string

//go:embed restore_token/clear.sql
var TokenClear string

//go:embed restore_token/delete.sql
var TokenDelete string

//go:embed credentials/get_by_login.sql
var AuthGetByLogin string

//go:embed credentials/get_by_id.sql
var AuthGetById string

//go:embed credentials/insert.sql
var AuthInsert string

//go:embed credentials/update.sql
var AuthUpdate string

//go:embed credentials/delete.sql
var AuthDelete string
