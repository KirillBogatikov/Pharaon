package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-auth/data/sql"
)

type HistoryRepository struct {
	repo *psql.SqlRepository
}

func NewHistoryRepository() (*HistoryRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &HistoryRepository{repo}, nil
}

func (h *HistoryRepository) Get(authId uuid.UUID) (*History, error) {
	rows, err := h.repo.DB.Queryx(sql.HistoryGet)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	history := &History{}
	err = rows.StructScan(history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (h *HistoryRepository) List(authId uuid.UUID) ([]History, error) {
	rows, err := h.repo.DB.Queryx(sql.HistoryList, authId)
	if err != nil {
		return nil, err
	}

	result := make([]History, 0)
	for rows.Next() {
		h := &History{}
		err := rows.StructScan(&h)

		if err != nil {
			return nil, err
		}

		result = append(result, *h)
	}

	return result, nil
}

func (h *HistoryRepository) Insert(history History) error {
	_, err := h.repo.DB.NamedExec(sql.HistoryInsert, history)
	return err
}

func (h *HistoryRepository) Delete(id uuid.UUID) (bool, error) {
	result, err := h.repo.DB.Exec(sql.HistoryDelete, id)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
