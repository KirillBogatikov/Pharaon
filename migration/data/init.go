package data

import (
	"fmt"
	config "github.com/Projector-Solutions/Pharaon-config/migration"
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"pharaon-migration/data/sql"
	"strings"
)

type InitRepository struct {
	repo *psql.SqlRepository
}

func NewInitRepository() (*InitRepository, error) {
	repo, err := psql.NewSqlRepository(*config.Database)
	if err != nil {
		return nil, err
	}

	return &InitRepository{repo}, nil
}

func (m *InitRepository) Check() (bool, error) {
	row := m.repo.DB.QueryRowx(sql.InitCheck)

	exists := false
	err := row.Scan(&exists)
	return exists, err
}

func (m *InitRepository) Create() error {
	_, err := m.repo.DB.Exec(sql.InitCreate)
	return err
}

func (m *InitRepository) InitService(service, user, password string) error {
	script := fmt.Sprintf(sql.InitService, user, password, service, user)

	for _, line := range strings.Split(script, ";") {
		_, err := m.repo.DB.Exec(line)
		if err != nil {
			return err
		}
	}

	return nil
}
