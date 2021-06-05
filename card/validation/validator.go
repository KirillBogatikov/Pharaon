package validation

import (
	"github.com/Projector-Solutions/Pharaon-api/tool"
	"github.com/Projector-Solutions/Pharaon-api/user"
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"github.com/google/uuid"
	"net/http"
	"pharaon-card/model"
)

var (
	v = &validation.Validator{}
)

func ValidateUsers(token string, userIds ...uuid.UUID) ([]validation.FieldResult, error) {
	status, exists, err := user.Check(userIds, token)
	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, tool.UnexpectedStatusError
	}

	result := make([]validation.FieldResult, len(exists))
	for i, r := range exists {
		if r {
			result[i] = validation.Valid
		} else {
			result[i] = NotFound
		}
	}

	return result, nil
}

func ValidateTag(tag *model.Tag) *TagResult {
	return &TagResult{
		Name: v.ValidateString(tag.Name, TagNameRule),
	}
}

func ValidateCard(card *model.Card, token string) (*CardResult, error) {
	result, err := ValidateUsers(token, *card.Author.Id)
	if err != nil {
		return nil, err
	}

	authorResult := result[0]
	var executorResult *validation.FieldResult = nil

	if card.Executor != nil {
		result, err = ValidateUsers(token, *card.Executor.Id)
		if err != nil {
			return nil, err
		}

		executorResult = &result[0]
	}

	cardResult := &CardResult{
		Name:        v.ValidateString(card.Name, NameRule),
		Description: v.ValidateString(card.Description, DescriptionRule),
		Author:      authorResult,
		Executor:    executorResult,
		Type:        v.ValidateInt(int64(card.Type), TypeRule),
		Priority:    v.ValidateInt(int64(card.Priority), PriorityRule),
	}

	observerIds := make([]uuid.UUID, len(card.Observers))
	for i, u := range card.Observers {
		observerIds[i] = *u.Id
	}

	cardResult.Observers, err = ValidateUsers(token, observerIds...)
	if err != nil {
		return nil, err
	}

	cardResult.Tags = make([]TagResult, len(card.Tags))
	for i, tag := range card.Tags {
		cardResult.Tags[i] = *ValidateTag(&tag)
	}

	return cardResult, nil
}
