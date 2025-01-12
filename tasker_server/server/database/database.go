package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type DatabaseImpl interface {
	GetMembershipsForUser(ctx context.Context, userID uint64) ([]Member, error)
	GetTodoListMembers(ctx context.Context, todoListID uint64) ([]Member, error)
	GetTodoEvents(ctx context.Context, todoID uint64) ([]TodoEvent, error)
	CreateTodoList(ctx context.Context, name string, token uuid.UUID, user *User) (uint64, error)
	GetTodoList(ctx context.Context, query *TodoListQueryParam) (*TodoList, error)
	ListTodoLists(ctx context.Context, userID uint64) ([]TodoList, error)
	JoinTodoList(ctx context.Context, todoListID uint64, userID uint64) error
	GetTodos(ctx context.Context, todoListID uint64) ([]Todo, error)
	GetTodo(ctx context.Context, todoID uint64) (*Todo, error)
	UpdateTodo(ctx context.Context, todo *Todo, user *User, nextStatus TodoStatus) error
	CreateTodo(ctx context.Context, todo *Todo) (uint64, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id uint64) (*User, error)
	CreateUser(ctx context.Context, email string) (uint64, error)
}

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) DatabaseImpl {
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
