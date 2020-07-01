package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TxFn func(context.Context, *sqlx.Tx) error

func WithTransaction(ctx context.Context, db *sqlx.DB, fn TxFn) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(ctx, tx)
	return err
}
