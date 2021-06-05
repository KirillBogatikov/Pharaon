package data

import (
	config "github.com/Projector-Solutions/Pharaon-config/user"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-user/auth"
	"pharaon-user/data/sql"
	"pharaon-user/model"
	"pharaon-user/personal"
)

type Repository struct {
	repo *psql.SqlRepository
}

func NewRepository() (*Repository, error) {
	repo, err := psql.NewSqlRepository(config.Database)
	if err != nil {
		return nil, err
	}

	return &Repository{repo}, nil
}

func (r *Repository) Check(ids ...uuid.UUID) ([]bool, error) {
	rows, err := r.repo.DB.Queryx(sql.UserCheckExists, UserIdList{Ids: ids})
	if err != nil {
		return nil, err
	}

	existIds := make([]uuid.UUID, 0)
	err = rows.StructScan(&existIds)
	if err != nil {
		return nil, err
	}

	mapped := make(map[uuid.UUID]bool)
	for _, id := range existIds {
		mapped[id] = true
	}

	result := make([]bool, len(ids))
	for i, id := range ids {
		_, ok := mapped[id]
		result[i] = ok
	}

	return result, nil
}

func (r *Repository) GetById(id uuid.UUID) (*model.User, error) {
	dbUser := &DBUser{Id: id}
	rows, err := r.repo.DB.Queryx(sql.UserGetById, dbUser)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	err = rows.StructScan(dbUser)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id: &dbUser.Id,
		Credentials: &auth.Credentials{
			Id: &dbUser.CredentialsId,
		},
		Personal: &personal.Data{
			Id: &dbUser.PersonalDataId,
		},
	}, nil
}

func (r *Repository) Insert(user *model.User) error {
	_, err := r.repo.DB.Exec(sql.UserInsert, &DBUser{
		Id:             *user.Id,
		CredentialsId:  *user.Credentials.Id,
		PersonalDataId: *user.Personal.Id,
	})
	return err
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	result, err := r.repo.DB.Exec(sql.UserDelete, &DBUser{Id: id})
	if err != nil {
		return false, err
	}

	count, err := result.RowsAffected()
	return count > 0, err
}
