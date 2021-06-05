package api

import (
	"github.com/google/uuid"
)

type ServiceData struct {
	Name string `json:"name"`
}

type Service struct {
	Id      *uuid.UUID `db:"id" json:"id"`
	Name    string     `db:"name" json:"name"`
	Version int        `db:"version" json:"version"`
}

type Migration struct {
	Version int    `json:"-"`
	Up      string `json:"up"`
	Down    string `json:"down"`
}

type MigrationResult struct {
	Service string  `json:"service"`
	Version int     `json:"version"`
	Error   *string `json:"error"`
}
