package service

import (
	"github.com/google/uuid"
	"pharaon-personal/model"
	"pharaon-personal/validation"
)

func GetById(id uuid.UUID) (*model.Data, error) {
	data, err := repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Signup(data *model.Data) (*validation.DataResult, error) {
	result, err := validation.ValidateData(repo, data)
	if err != nil {
		return nil, err
	}

	if !result.IsValid() {
		return result, nil
	}

	dataId := uuid.New()
	data.Id = &dataId

	nameId := uuid.New()
	data.Name.Id = &nameId

	err = repo.Insert(data)
	return result, err
}

func Merge(data *model.Data) (bool, error) {
	oldData, err := repo.GetById(*data.Id)
	if err != nil {
		return false, err
	}

	if oldData == nil {
		return false, nil
	}

	if data.BirthDate == nil {
		data.BirthDate = oldData.BirthDate
	}

	if len(data.Phone) == 0 {
		data.Phone = oldData.Phone
	}

	oldName := oldData.Name
	newName := data.Name

	if newName == nil {
		data.Name = oldName
	} else {
		if len(newName.First) == 0 {
			newName.First = oldName.First
		}

		if len(newName.Last) == 0 {
			newName.Last = oldName.Last
		}
	}

	return true, nil
}

func Update(data *model.Data) (bool, *validation.DataResult, error) {
	if data.Id == nil {
		return false, nil, nil
	}

	found, err := Merge(data)
	if err != nil {
		return false, nil, err
	}

	if !found {
		return false, nil, nil
	}

	result, err := validation.ValidateData(repo, data)
	if err != nil {
		return false, nil, err
	}

	if !result.IsValid() {
		return false, result, nil
	}

	found, err = repo.Update(data)
	if err != nil {
		return false, result, err
	}

	return found, result, nil
}

func Delete(id uuid.UUID) (bool, error) {
	return repo.Delete(id)
}

func Validate(data *model.Data) (bool, *validation.DataResult, error) {
	if data.Id != nil {
		found, err := Merge(data)
		if err != nil {
			return false, nil, err
		}

		if !found {
			return false, nil, nil
		}
	}

	result, err := validation.ValidateData(repo, data)
	return true, result, err
}
