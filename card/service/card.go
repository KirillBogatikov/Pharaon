package service

import (
	"github.com/Projector-Solutions/Pharaon-api/user"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"pharaon-card/model"
	"pharaon-card/validation"
)

func GetCards(token string, ids ...uuid.UUID) ([]model.Card, error) {
	cards, err := cardRepo.GetAll(ids)
	if err != nil {
		return nil, err
	}

	result := make([]model.Card, 0)
	for _, card := range cards {
		skip, err := cleanCard(token, &card)
		if err != nil {
			return nil, err
		}
		if skip {
			continue
		}

		result = append(result, card)
	}

	return result, nil
}

func GetCard(token string, id uuid.UUID) (*model.Card, error) {
	card, err := cardRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	skip, err := cleanCard(token, card)
	if err != nil {
		return nil, err
	}

	if skip {
		return nil, nil
	}

	return card, nil
}

func Merge(token string, card *model.Card) (bool, error) {
	oldCard, err := GetCard(token, *card.Id)
	if err != nil {
		return false, err
	}

	if oldCard == nil {
		return false, nil
	}

	if len(card.Name) == 0 {
		card.Name = oldCard.Name
	}

	if card.Author.Id == nil {
		card.Author.Id = oldCard.Author.Id
	}

	if card.Type == 0 {
		card.Type = oldCard.Type
	}

	if card.Priority == 0 {
		card.Priority = oldCard.Priority
	}

	if card.Observers == nil {
		card.Observers = make([]user.User, 0)
	}

	if card.Tags == nil {
		card.Tags = make([]model.Tag, 0)
	}

	return true, nil
}

func SaveCard(token string, card *model.Card) (found bool, result *validation.CardResult, err error) {
	if card.Id != nil {
		found, err = Merge(token, card)
		if err != nil {
			return
		}

		if !found {
			return
		}
	}

	result, err = validation.ValidateCard(card, token)
	if err != nil {
		return
	}

	if !result.IsValid() {
		return
	}

	var tx *sqlx.Tx

	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	if card.Id == nil {
		id := uuid.New()
		card.Id = &id

		tx, err = cardRepo.Insert(card)

		if err != nil {
			return
		}
	} else {
		tx, found, err = cardRepo.Update(card)

		if err != nil {
			return
		}

		if !found {
			return
		}

		_, err = cardTagRepo.Clear(tx, *card.Id)
		if err != nil {
			return
		}

		_, err = observerRepo.Clear(tx, *card.Id)
		if err != nil {
			return
		}
	}

	for _, tag := range card.Tags {
		err = cardTagRepo.Insert(tx, uuid.New(), *card.Id, *tag.Id)
		if err != nil {
			return
		}
	}

	for _, u := range card.Observers {
		err = observerRepo.Insert(tx, uuid.New(), *card.Id, *u.Id)
		if err != nil {
			return
		}
	}

	return
}

func DeleteCard(id uuid.UUID) (bool, error) {
	return cardRepo.Delete(id)
}
