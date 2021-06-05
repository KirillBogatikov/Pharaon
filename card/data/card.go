package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/user"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"pharaon-card/data/sql"
	"pharaon-card/model"
)

type CardRepository struct {
	repo *psql.SqlRepository
}

func NewCardRepository() (*CardRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &CardRepository{repo}, nil
}

func (c *CardRepository) GetById(id uuid.UUID) (*model.Card, error) {
	rows, err := c.repo.DB.Queryx(sql.CardGetById, &DBCard{Id: id})
	if err != nil {
		return nil, err
	}

	array, err := c.parse(rows)
	if err != nil {
		return nil, err
	}

	if len(array) == 0 {
		return nil, nil
	}

	return &array[0], nil
}

func (c *CardRepository) GetAll(ids []uuid.UUID) ([]model.Card, error) {
	rows, err := c.repo.DB.Queryx(sql.CardGetInIds, &CardIdList{Ids: ids})
	if err != nil {
		return nil, err
	}

	array, err := c.parse(rows)
	return array, err
}

func (c *CardRepository) parse(rows *sqlx.Rows) ([]model.Card, error) {
	dbCards := make([]DBCard, 0)

	err := rows.StructScan(&dbCards)
	if err != nil {
		return nil, err
	}

	cards := make([]model.Card, len(dbCards))
	for i, dbCard := range dbCards {
		cards[i] = *DBToCard(&dbCard)
	}

	return cards, nil
}

func (c *CardRepository) Insert(card *model.Card) (*sqlx.Tx, error) {
	tx, err := c.repo.DB.Beginx()
	if err != nil {
		return nil, err
	}

	_, err = tx.NamedExec(sql.CardInsert, CardToDB(card))
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return tx, nil
}

func (c *CardRepository) Update(card *model.Card) (*sqlx.Tx, bool, error) {
	tx, err := c.repo.DB.Beginx()
	if err != nil {
		return nil, false, err
	}

	result, err := tx.NamedExec(sql.CardUpdate, CardToDB(card))
	if err != nil {
		return nil, false, err
	}

	count, err := result.RowsAffected()
	return tx, count > 0, err
}

func (c *CardRepository) Delete(id uuid.UUID) (bool, error) {
	result, err := c.repo.DB.Exec(sql.CardDelete, &DBCard{Id: id})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
