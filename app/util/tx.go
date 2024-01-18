package util

import (
	"context"
	"database/sql"
)

func ExecTx(ctx context.Context, dbs *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := dbs.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
