package service

import "pharaon-card/data"

var (
	cardRepo     *data.CardRepository
	tagRepo      *data.TagRepository
	observerRepo *data.ObserverRepository
	cardTagRepo  *data.CardTagRepository
)

func InitRepository() error {
	var err error

	cardRepo, err = data.NewCardRepository()
	if err != nil {
		return err
	}

	tagRepo, err = data.NewTagRepository()
	if err != nil {
		return err
	}

	observerRepo, err = data.NewObserverRepository()
	if err != nil {
		return err
	}

	cardTagRepo, err = data.NewCardTagRepository()
	if err != nil {
		return err
	}

	return nil
}
