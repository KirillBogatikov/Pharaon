package data

import "github.com/google/uuid"

type DBCard struct {
	Id          uuid.UUID  `db:"card_id"`
	Name        string     `db:"card_name"`
	Description string     `db:"description"`
	AuthorId    uuid.UUID  `db:"author_id"`
	ExecutorId  *uuid.UUID `db:"executor_id"`
	Type        int        `db:"type"`
	Priority    int        `db:"priority"`
}

type DBObserver struct {
	Id     uuid.UUID `db:"observer_id"`
	CardId uuid.UUID `db:"card_id"`
	UserId uuid.UUID `db:"user_id"`
}

type DBTag struct {
	TagId uuid.UUID `db:"tag_id"`
	Name  string    `db:"tag_name"`
}

type CardTag struct {
	Id     uuid.UUID `db:"card_tag_id"`
	CardId uuid.UUID `db:"card_id"`
	*DBTag
}

type CardIdList struct {
	Ids []uuid.UUID `db:"card_id_list"`
}
