package data

import (
	_ "embed"
	config "github.com/Projector-Solutions/Pharaon-config/migration"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"pharaon-migration/data/sql"
)

type ServiceRepository struct {
	repo *psql.SqlRepository
}

func NewServiceRepository() (*ServiceRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &ServiceRepository{repo}, nil
}

func (r *ServiceRepository) GetByName(name string) (*Service, error) {
	row, err := r.repo.DB.Queryx(sql.ServiceGetByName, name)

	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, nil
	}

	service := Service{}
	err = row.StructScan(&service)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *ServiceRepository) Insert(s *Service) error {
	_, err := r.repo.DB.NamedExec(sql.ServiceInsert, s)
	return err
}

func (r *ServiceRepository) UpdateVersion(name string, version int) error {
	_, err := r.repo.DB.Exec(sql.ServiceUpdate, version, name)
	return err
}

func (r *ServiceRepository) Delete(name string) error {
	_, err := r.repo.DB.Exec(sql.ServiceDelete, name)
	return err
}
