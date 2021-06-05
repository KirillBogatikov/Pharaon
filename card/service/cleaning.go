package service

import (
	"github.com/Projector-Solutions/Pharaon-api/user"
	"pharaon-card/api"
	"pharaon-card/model"
)

func cleanObservers(token string, card *model.Card) error {
	result := make([]user.User, 0)
	for _, observer := range card.Observers {
		u, err := api.GetUserById(*observer.Id, token)
		if err != nil {
			return err
		}

		if u == nil {
			go func() {
				_, _ = observerRepo.Delete(*observer.Id)
			}()
			continue
		}

		result = append(result, *u)
	}

	card.Observers = result
	return nil
}

func cleanCard(token string, card *model.Card) (bool, error) {
	u, err := api.GetUserById(*card.Author.Id, token)
	if err != nil {
		return false, err
	}

	if u == nil {
		go func() {
			_, _ = cardRepo.Delete(*card.Id)
		}()
		return true, nil
	} else {
		card.Author = *u
	}

	if card.Executor != nil {
		u, err = api.GetUserById(*card.Author.Id, token)
		if err != nil {
			return false, err
		}

		card.Executor = u
	}

	return false, cleanObservers(token, card)
}
