package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-auth/data/sql"
	"time"
)

type TokenRepository struct {
	repo *psql.SqlRepository
}

func NewTokenRepository() (*TokenRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &TokenRepository{repo}, nil
}

func (t *TokenRepository) get(sql string, value interface{}) (*RestoreToken, error) {
	rows, err := t.repo.DB.Queryx(sql, value)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	token := &RestoreToken{}
	err = rows.StructScan(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *TokenRepository) GetByAuth(authId uuid.UUID) (*RestoreToken, error) {
	return t.get(sql.TokenGetByAuth, authId)
}

func (t *TokenRepository) GetByToken(token string) (*RestoreToken, error) {
	return t.get(sql.TokenGetByToken, token)
}

func (t *TokenRepository) Insert(token RestoreToken) error {
	_, err := t.repo.DB.NamedExec(sql.TokenInsert, token)
	return err
}

func (t *TokenRepository) Clear() (int, error) {
	res, err := t.repo.DB.Exec(sql.TokenClear, time.Now())
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(count), err
}

func (t *TokenRepository) Delete(id uuid.UUID) (bool, error) {
	res, err := t.repo.DB.Exec(sql.TokenDelete, id)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	return count > 0, err
}
