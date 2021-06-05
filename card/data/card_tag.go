package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/card"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"pharaon-card/data/sql"
	"pharaon-card/model"
)

type CardTagRepository struct {
	repo *psql.SqlRepository
}

func NewCardTagRepository() (*CardTagRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &CardTagRepository{repo}, nil
}

func (ct *CardTagRepository) GetForCards(cards ...uuid.UUID) (map[uuid.UUID][]model.Tag, error) {
	result := make(map[uuid.UUID][]model.Tag, 0)

	rows, err := ct.repo.DB.Queryx(sql.TagGetForCards, &CardIdList{Ids: cards})
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tag := &CardTag{}
		err = rows.StructScan(tag)
		if err != nil {
			return nil, err
		}

		array, ok := result[tag.CardId]
		if !ok {
			array = make([]model.Tag, 0)
		}

		result[tag.CardId] = append(array, model.Tag{
			Id:   &tag.Id,
			Name: tag.Name,
		})
	}

	return result, nil
}

func (ct *CardTagRepository) Insert(tx *sqlx.Tx, id, cardId, tagId uuid.UUID) error {
	_, err := tx.NamedExec(sql.CardTagInsert, CardTag{
		Id:     id,
		CardId: cardId,
		DBTag:  &DBTag{TagId: tagId},
	})

	return err
}

func (ct *CardTagRepository) Clear(tx *sqlx.Tx, cardId uuid.UUID) (bool, error) {
	result, err := tx.NamedExec(sql.CardTagClear, CardTag{
		CardId: cardId,
	})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
