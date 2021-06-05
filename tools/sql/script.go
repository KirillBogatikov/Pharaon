package psql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type empty struct{}

func (s *SqlRepository) NamedExecMany(script string, items ...interface{}) (result []sql.Result, err error) {
	tx, err := s.DB.Beginx()
	if err != nil {
		return nil, err
	}

	var success bool
	success, result, err = s.NamedExecTx(tx, script, items...)

	if success {
		err = tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}

func (s *SqlRepository) NamedExecManyContext(ctx context.Context, script string, items ...interface{}) (result []sql.Result, err error) {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var success bool
	success, result, err = s.NamedExecTx(tx, script, items...)

	if success {
		err = tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}

func (s *SqlRepository) NamedExecTx(tx *sqlx.Tx, script string, items ...interface{}) (txSuccess bool, result []sql.Result, err error) {
	txSuccess = true

	for i, line := range strings.Split(script, ";") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		var item interface{} = empty{}

		if len(items) > i && items[i] != nil {
			item = items[i]
		}

		var res sql.Result
		res, err = tx.NamedExec(line, item)
		if err != nil {
			log.Printf("query #%d failed: %s\n", i, err)
			txSuccess = false
			break
		}

		result = append(result, res)
	}

	return
}
