package sql

import _ "embed"

//go:embed observer/get_for_cards.sql
var ObserverGetForCards string

//go:embed observer/insert.sql
var ObserverInsert string

//go:embed observer/clear.sql
var ObserverClear string

//go:embed observer/delete.sql
var ObserverDelete string
