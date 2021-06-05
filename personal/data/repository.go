package data

import (
	sql2 "database/sql"
	config "github.com/Projector-Solutions/Pharaon-config/personal"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/google/uuid"
	"pharaon-personal/data/sql"
	"pharaon-personal/model"
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

func (r *Repository) CheckPhoneRegistered(phone string) (*uuid.UUID, error) {
	rows, err := r.repo.DB.NamedQuery(sql.DataCheckPhone, &model.Data{Phone: phone})
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	var id uuid.UUID
	err = rows.Scan(&id)
	return &id, err
}

func (r *Repository) GetById(id uuid.UUID) (*model.Data, error) {
	rows, err := r.repo.DB.NamedQuery(sql.DataGetById, &model.Data{Id: &id})
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	ndata := &NamedData{}
	err = rows.StructScan(ndata)
	if err != nil {
		return nil, err
	}

	ndata.Data.Name = &ndata.Name
	return &ndata.Data, nil
}

func (r *Repository) Insert(data *model.Data) error {
	ndata := NewNamedData(data, data.Name)
	_, error := r.repo.NamedExecMany(sql.DataInsert, ndata, ndata)
	return error
}

func (r *Repository) found(results ...sql2.Result) bool {
	for _, result := range results {
		c, _ := result.RowsAffected()
		if c == 0 {
			return false
		}
	}

	return true
}

func (r *Repository) Update(data *model.Data) (bool, error) {
	ndata := NewNamedData(data, data.Name)
	results, err := r.repo.NamedExecMany(sql.DataUpdate, ndata, ndata)
	return r.found(results...), err
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	results, err := r.repo.DB.NamedExec(sql.DataDelete, &model.Data{Id: &id})
	return r.found(results), err
}
