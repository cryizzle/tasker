package database

import (
	"context"

	"github.com/judedaryl/go-arrayutils"
)

type User struct {
	ID          uint64   `json:"user_id" db:"user_id"`
	Email       string   `json:"email" db:"email"`
	Memberships []Member `json:"memberships"`
}

func (u User) IsMember(todoListID uint64) bool {
	member := arrayutils.Find(u.Memberships, func(member Member) bool {
		return member.TodoListID == todoListID
	})
	return member != nil
}

func (db Database) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT * FROM users	WHERE email = ?`

	err := db.db.GetContext(ctx, &user, query, email)

	if err != nil {
		return nil, err
	}

	members, err := db.GetMembershipsForUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	user.Memberships = members

	return &user, nil
}

func (db Database) GetUserByID(ctx context.Context, id uint64) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE user_id = ?`

	err := db.db.GetContext(ctx, &user, query, id)

	if err != nil {
		return nil, err
	}

	members, err := db.GetMembershipsForUser(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Memberships = members

	return &user, nil
}

func (db Database) CreateUser(ctx context.Context, email string) (uint64, error) {
	query := `INSERT IGNORE INTO users (email) VALUES (?)`
	result, err := db.db.ExecContext(ctx, query, email)

	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}
