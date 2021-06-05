package data

import (
	psql "github.com/Projector-Solutions/Pharaon-tools/sql"
	"github.com/jmoiron/sqlx"
)

type MigrationRepository struct {
	repo *psql.SqlRepository
}

func NewMigrationRepository(config *psql.DatabaseConfig) (*MigrationRepository, error) {
	repo, err := psql.NewSqlRepository(*config)
	if err != nil {
		return nil, err
	}

	return &MigrationRepository{repo}, nil
}

func (m *MigrationRepository) Execute(script string) (*sqlx.Tx, error) {
	tx, err := m.repo.DB.Beginx()
	if err != nil {
		return nil, err
	}

	_, _, err = m.repo.NamedExecTx(tx, script)
	return tx, err
}

func (m *MigrationRepository) Close() {
	_ = m.repo.Close()
}
