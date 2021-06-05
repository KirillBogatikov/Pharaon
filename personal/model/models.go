package model

import (
	"github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
)

type Data struct {
	Id        *uuid.UUID `db:"data_id" json:"id"`
	Phone     string     `db:"phone" json:"phone"`
	BirthDate *time.Date `db:"birth_date" json:"birth_date"`
	Photo     *uuid.UUID `db:"photo_id" json:"photo_id,omitempty"`
	Name      *Name      `json:"name"`
}

type Name struct {
	Id         *uuid.UUID `db:"name_id" json:"-"`
	DataId     *uuid.UUID `db:"data_id" json:"-"`
	First      string     `db:"first" json:"first"`
	Last       string     `db:"last" json:"last"`
	Patronymic string     `db:"patronymic" json:"patronymic"`
}
