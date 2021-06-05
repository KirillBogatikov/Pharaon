package personal

import (
	config "github.com/Projector-Solutions/Pharaon-config/personal"
	"github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"github.com/google/uuid"
)

var (
	personalApiURL = ""
)

func init() {
	if personalApiURL = config.Http.ApiUrl; len(personalApiURL) == 0 {
		panic("personal data service URL required")
	}
}

type Data struct {
	Id        *uuid.UUID `json:"id"`
	Phone     string     `json:"phone"`
	BirthDate *time.Date `json:"birth_date"`
	Photo     *uuid.UUID `json:"photo_id,omitempty"`
	Name      *Name      `json:"name"`
}

type Name struct {
	First      string `json:"first"`
	Last       string `json:"last"`
	Patronymic string `json:"patronymic"`
}

type NameResult struct {
	First      validation.FieldResult `json:"first"`
	Last       validation.FieldResult `json:"last"`
	Patronymic validation.FieldResult `json:"patronymic"`
}

func (m *NameResult) IsValid() bool {
	return m.First == validation.Valid &&
		m.Last == validation.Valid &&
		m.Patronymic == validation.Valid
}

type DataResult struct {
	Phone     validation.FieldResult `json:"phone"`
	BirthDate validation.FieldResult `json:"birth_date"`
	Name      *NameResult            `json:"name"`
}

func (m *DataResult) IsValid() bool {
	return m.Phone == validation.Valid && m.BirthDate == validation.Valid && m.Name.IsValid()
}
