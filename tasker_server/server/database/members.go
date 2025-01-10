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
	Membership Membership `json:"type" db:"membership"`
	User
}

func (db Database) GetMembershipsForUser(ctx context.Context, userID uint64) ([]Member, error) {
	var members []Member
	query := `SELECT
		members.todo_list_id as todo_list_id,
		members.membership as membership,
		users.user_id as user_id,
		users.email as email
	 	FROM members
		JOIN users ON members.user_id = users.user_id
		WHERE members.user_id = ?`

	err := db.db.SelectContext(ctx, &members, query, userID)
	return members, err
}

func (db Database) GetTodoListMembers(ctx context.Context, todoListID uint64) ([]Member, error) {
	var members []Member
	query := `SELECT
		members.todo_list_id as todo_list_id,
		members.membership as membership,
		users.user_id as user_id,
		users.email as email
	 	FROM members
		JOIN users ON members.user_id = users.user_id
		WHERE members.todo_list_id = ?`

	err := db.db.SelectContext(ctx, &members, query, todoListID)
	return members, err
}

func createMember(ctx context.Context, tx *sqlx.Tx, todoListID uint64, userID uint64, membership Membership) error {
	query := `INSERT INTO members (todo_list_id, user_id, membership) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, todoListID, userID, membership)

	return err
}
