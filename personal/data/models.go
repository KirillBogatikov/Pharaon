package data

import "pharaon-personal/model"

type NamedData struct {
	model.Data
	model.Name
}

func NewNamedData(data *model.Data, name *model.Name) *NamedData {
	return &NamedData{*data, *name}
}
