package service

import (
	"fmt"
	"github.com/google/uuid"
	"pharaon-card/model"
	"pharaon-card/validation"
)

func SearchTag(query string) ([]model.Tag, error) {
	return tagRepo.Search(fmt.Sprintf("^.*%s.*$", query))
}

func ListTags() ([]model.Tag, error) {
	return tagRepo.List()
}

func InsertTag(tag *model.Tag) (*validation.TagResult, error) {
	result := validation.ValidateTag(tag)
	if !result.IsValid() {
		return result, nil
	}

	id := uuid.New()
	tag.Id = &id
	err := tagRepo.Insert(tag)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func DeleteTag(id uuid.UUID) (bool, error) {
	return tagRepo.Delete(id)
}
