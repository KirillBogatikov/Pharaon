package validation

import (
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"pharaon-card/model"
)

const (
	NotFound validation.FieldResult = 5
)

type TagResult struct {
	Name validation.FieldResult
}

func (t *TagResult) IsValid() bool {
	return t.Name == validation.Valid
}

type CardResult struct {
	Name        validation.FieldResult   `json:"name"`
	Description validation.FieldResult   `json:"description"`
	Author      validation.FieldResult   `json:"author"`
	Executor    *validation.FieldResult  `json:"executor,omitempty"`
	Type        validation.FieldResult   `json:"type"`
	Priority    validation.FieldResult   `json:"priority"`
	Observers   []validation.FieldResult `json:"observers"`
	Tags        []TagResult              `json:"tags"`
}

func (r *CardResult) IsValid() bool {
	observersValid := true
	for _, o := range r.Observers {
		if o != validation.Valid {
			observersValid = false
			break
		}
	}

	tagsValid := true
	for _, t := range r.Tags {
		if !t.IsValid() {
			tagsValid = false
			break
		}
	}

	return r.Name == validation.Valid && r.Description == validation.Valid &&
		r.Author == validation.Valid && (r.Executor == nil || *r.Executor == validation.Valid) &&
		r.Type == validation.Valid && r.Priority == validation.Valid &&
		observersValid && tagsValid
}

var NameRule = validation.Rule{
	Min:   4,
	Max:   1024,
	Regex: ".+",
}

var DescriptionRule = validation.Rule{
	Min:   0,
	Max:   10000,
	Regex: ".+",
}

var TypeRule = validation.Rule{
	Min: int64(model.TypeTask),
	Max: int64(model.TypeProposal),
}

var PriorityRule = validation.Rule{
	Min: int64(model.PriorityLow),
	Max: int64(model.PriorityTop),
}

var TagNameRule = validation.Rule{
	Min:   2,
	Max:   128,
	Regex: ".+",
}
