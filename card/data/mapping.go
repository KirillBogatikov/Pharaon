package data

import (
	"github.com/Projector-Solutions/Pharaon-api/user"
	"github.com/google/uuid"
	"pharaon-card/model"
)

func CardToDB(card *model.Card) *DBCard {
	var executorId *uuid.UUID = nil
	if card.Executor != nil {
		executorId = card.Executor.Id
	}

	return &DBCard{
		Id:          *card.Id,
		Name:        card.Name,
		Description: card.Description,
		AuthorId:    *card.Author.Id,
		ExecutorId:  executorId,
		Type:        int(card.Type),
	}
}

func DBToCard(dbCard *DBCard) *model.Card {
	var executor *user.User = nil
	if dbCard.ExecutorId != nil {
		executor = &user.User{Id: dbCard.ExecutorId}
	}

	return &model.Card{
		Id:          &dbCard.Id,
		Name:        dbCard.Name,
		Description: dbCard.Description,
		Author:      user.User{Id: &dbCard.AuthorId},
		Executor:    executor,
		Type:        model.Type(dbCard.Type),
		Priority:    model.Priority(dbCard.Priority),
		Observers:   nil,
		Tags:        nil,
	}
}
