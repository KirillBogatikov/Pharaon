package validation

import (
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"pharaon-personal/data"
	"pharaon-personal/model"
)

var v = &validation.Validator{}

func ValidateName(name *model.Name) *NameResult {
	if name == nil {
		return &NameResult{validation.Incorrect, validation.Incorrect, validation.Incorrect}
	}

	return &NameResult{
		First:      v.ValidateString(name.First, FirstNameRule),
		Last:       v.ValidateString(name.Last, LastNameRule),
		Patronymic: v.ValidateString(name.Patronymic, PatronymicRule),
	}
}

func ValidateData(repo *data.Repository, data *model.Data) (*DataResult, error) {
	birthDateResult := validation.Valid

	if data.BirthDate == nil {
		birthDateResult = validation.Incorrect
	} else if data.BirthDate.Before(MinBirthDate) {
		birthDateResult = validation.Long
	} else if data.BirthDate.After(MaxBirthDate) {
		birthDateResult = validation.Short
	}

	result := DataResult{
		Phone:     v.ValidateString(data.Phone, PhoneRule),
		BirthDate: birthDateResult,
		Name:      ValidateName(data.Name),
	}

	phoneOwnerId, err := repo.CheckPhoneRegistered(data.Phone)
	if err != nil {
		return nil, err
	}

	if phoneOwnerId != nil {
		if data.Id == nil {
			result.Phone = validation.Busy
		} else if phoneOwnerId.String() != data.Id.String() {
			result.Phone = validation.Busy
		}
	}

	return &result, nil
}
