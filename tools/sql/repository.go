package psql

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var connections = make(map[DatabaseConfig]*sqlx.DB)

type SqlRepository struct {
	config DatabaseConfig
	DB     *sqlx.DB
}

func NewSqlRepository(dbConfig DatabaseConfig) (repo *SqlRepository, err error) {
	db, exists := connections[dbConfig]

	if !exists {
		repo, err = ForceSqlRepository(dbConfig)
		if err != nil {
			return nil, err
		}

		connections[dbConfig] = repo.DB
		return
	}

	return &SqlRepository{dbConfig, db}, nil
}

func ForceSqlRepository(dbConfig DatabaseConfig) (*SqlRepository, error) {
	connConfig, err := pgx.ParseConfig(dbConfig.URL)
	if err != nil {
		return nil, err
	}

	nativeDB := stdlib.OpenDB(*connConfig)
	db := sqlx.NewDb(nativeDB, "pgx")
	return &SqlRepository{dbConfig, db}, nil
}

func (s *SqlRepository) Close() error {
	delete(connections, s.config)
	return s.DB.Close()
}
