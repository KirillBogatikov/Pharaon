package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/card"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"pharaon-card/data/sql"
)

type ObserverRepository struct {
	repo *psql.SqlRepository
}

func NewObserverRepository() (*ObserverRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &ObserverRepository{repo}, nil
}

func (o *ObserverRepository) GetForCard(ids []uuid.UUID) (map[uuid.UUID][]uuid.UUID, error) {
	rows, err := o.repo.DB.Queryx(sql.ObserverGetForCards, CardIdList{Ids: ids})
	if err != nil {
		return nil, err
	}

	observers := make([]DBObserver, 0)
	err = rows.StructScan(&observers)
	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID][]uuid.UUID)
	for _, observer := range observers {
		userIds, ok := result[observer.CardId]
		if !ok {
			userIds = make([]uuid.UUID, 0)
		}

		result[observer.CardId] = append(userIds, observer.UserId)
	}

	return result, nil
}

func (o *ObserverRepository) Insert(tx *sqlx.Tx, id, cardId, userId uuid.UUID) error {
	_, err := tx.NamedExec(sql.ObserverInsert, DBObserver{
		Id:     id,
		CardId: cardId,
		UserId: userId,
	})
	return err
}

func (o *ObserverRepository) Clear(tx *sqlx.Tx, cardId uuid.UUID) (bool, error) {
	result, err := tx.NamedExec(sql.ObserverClear, &DBObserver{CardId: cardId})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}

func (o *ObserverRepository) Delete(userId uuid.UUID) (bool, error) {
	result, err := o.repo.DB.NamedExec(sql.ObserverDelete, &DBObserver{UserId: userId})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
