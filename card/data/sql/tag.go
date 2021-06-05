package sql

import _ "embed"

//go:embed tag/search_by_name.sql
var TagSearchByName string

//go:embed tag/get_by_name.sql
var TagGetByName string

//go:embed tag/list.sql
var TagList string

//go:embed tag/insert.sql
var TagInsert string

//go:embed tag/delete.sql
var TagDelete string

//go:embed card_tag/get_for_cards.sql
var TagGetForCards string

//go:embed card_tag/insert.sql
var CardTagInsert string

//go:embed card_tag/clear.sql
var CardTagClear string
