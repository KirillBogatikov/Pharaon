package validation

import (
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"time"
)

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

var FirstNameRule = validation.Rule{
	Min:   1,
	Max:   256,
	Regex: "^[a-zA-Zа-яА-Я ]+$",
}

var LastNameRule = validation.Rule{
	Min:   1,
	Max:   256,
	Regex: "^[a-zA-Zа-яА-Я \\-]+$",
}

var PatronymicRule = validation.Rule{
	Min:   0,
	Max:   256,
	Regex: "^[a-zA-Zа-яА-Я]*$",
}

var PhoneRule = validation.Rule{
	Min:   4,
	Max:   256,
	Regex: "^(\\+|\\*)?[()0-9 \\-]+$",
}

var MinBirthDate = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
var MaxBirthDate = time.Now().Truncate(time.Hour * 8760 * 14)
