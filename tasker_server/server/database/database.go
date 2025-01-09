package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{db: db}
}

func WithTransaction(db *sqlx.DB, ctx context.Context, opts *sql.TxOptions, txFunc func(context.Context, *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, opts)

	if err != nil {
		return fmt.Errorf("failed to begin transaction, %w", err)
	}

	err = txFunc(ctx, tx)

	if err != nil {
		if e := tx.Rollback(); e != nil {
			return fmt.Errorf("failed to execute transaction, %w", err)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction, %w", err)
	}

	return nil
}

func WithTransactionRet[T any](db *sqlx.DB, ctx context.Context, opts *sql.TxOptions, txFunc func(context.Context, *sqlx.Tx) (T, error)) (T, error) {
	var zero T
	tx, err := db.BeginTxx(ctx, opts)

	if err != nil {
		return zero, fmt.Errorf("failed to begin transaction, %w", err)
	}

	ret, err := txFunc(ctx, tx)

	if err != nil {
		if e := tx.Rollback(); e != nil {
			return zero, fmt.Errorf("failed to execute transaction, %w", err)
		}

		return zero, err
	}

	if err := tx.Commit(); err != nil {
		return zero, fmt.Errorf("failed to commit transaction, %w", err)
	}

	return ret, nil
}
