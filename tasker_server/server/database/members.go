package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Membership string

const (
	OWNER  Membership = "OWNER"
	MEMBER Membership = "MEMBER"
)

type Member struct {
	TodoListID uint64     `json:"todo_list_id" db:"todo_list_id"`
	UserID     uint64     `json:"user_id" db:"user_id"`
	Membership Membership `json:"type" db:"membership"`
}

func (db Database) GetMembershipsForUser(ctx context.Context, userID uint64) ([]Member, error) {
	var members []Member
	query := `SELECT * FROM members WHERE user_id = ?`

	err := db.db.SelectContext(ctx, &members, query, userID)
	return members, err
}

func (db Database) GetMembers(ctx context.Context, todoListID uint64) ([]Member, error) {
	var members []Member
	query := `SELECT * FROM members WHERE todo_list_id = ?`

	err := db.db.SelectContext(ctx, &members, query, todoListID)
	return members, err
}

func createMember(ctx context.Context, tx *sqlx.Tx, todoListID uint64, userID uint64, membership Membership) error {
	query := `INSERT INTO members (todo_list_id, user_id, membership) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, todoListID, userID, membership)

	return err
}
