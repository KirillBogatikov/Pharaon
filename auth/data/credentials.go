package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/auth"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-auth/data/sql"
)

type CredentialsRepository struct {
	repo *psql.SqlRepository
}

func NewCredentialsRepository() (*CredentialsRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &CredentialsRepository{repo}, nil
}

func (a *CredentialsRepository) get(sql string, value interface{}) (*Credentials, error) {
	rows, err := a.repo.DB.Queryx(sql, value)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	auth := &Credentials{}
	err = rows.StructScan(auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *CredentialsRepository) GetByLogin(login string) (*Credentials, error) {
	return a.get(sql.AuthGetByLogin, login)
}

func (a *CredentialsRepository) GetById(id uuid.UUID) (*Credentials, error) {
	return a.get(sql.AuthGetById, id)
}

func (a *CredentialsRepository) Insert(auth Credentials) error {
	_, err := a.repo.DB.NamedExec(sql.AuthInsert, auth)
	return err
}

func (a *CredentialsRepository) Update(auth Credentials) (bool, error) {
	result, err := a.repo.DB.NamedExec(sql.AuthUpdate, auth)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}

func (a *CredentialsRepository) Delete(id uuid.UUID) (bool, error) {
	result, err := a.repo.DB.Exec(sql.AuthDelete, id)
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
