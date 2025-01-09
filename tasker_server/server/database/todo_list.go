package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TodoList struct {
	ID      uint64   `json:"id" db:"todo_list_id"`
	Name    string   `json:"name" db:"name"`
	Members []Member `json:"members"`
	Todos   []Todo   `json:"todos"`
}

func (db Database) CreateTodoList(ctx context.Context, name string, user *User) (uint64, error) {

	return WithTransactionRet[uint64](db.db, ctx, nil, func(ctx context.Context, tx *sqlx.Tx) (uint64, error) {

		query := `INSERT INTO todo_lists (name) VALUES (?)`
		result, err := tx.ExecContext(ctx, query, name)

		if err != nil {
			return 0, err
		}

		todoListID, err := result.LastInsertId()

		if err != nil {
			return 0, err
		}

		err = createMember(ctx, tx, uint64(todoListID), user.ID, OWNER)

		if err != nil {
			return 0, err
		}

		return uint64(todoListID), nil

	})

}

func (db Database) GetTodoList(ctx context.Context, todoListID uint64) (*TodoList, error) {
	var todoList TodoList
	query := `SELECT * FROM todo_lists WHERE todo_list_id = ?`

	err := db.db.GetContext(ctx, &todoList, query, todoListID)

	if err != nil {
		return nil, err
	}

	members, err := db.GetMembers(ctx, todoListID)

	if err != nil {
		return nil, err
	}
	todoList.Members = members

	todos, err := db.GetTodos(ctx, todoListID)
	if err != nil {
		return nil, err
	}
	todoList.Todos = todos

	return &todoList, nil
}

func (db Database) ListTodoLists(ctx context.Context, userID uint64) ([]TodoList, error) {
	var todoLists []TodoList
	query := `SELECT
	tdl.todo_list_id as todo_list_id,
	tdl.name as name
	FROM todo_lists tdl
	JOIN members ON tdl.todo_list_id = members.todo_list_id
	WHERE members.user_id = ?`

	err := db.db.SelectContext(ctx, &todoLists, query, userID)
	return todoLists, err
}

func (db Database) JoinTodoList(ctx context.Context, todoListID uint64, userID uint64) error {

	return WithTransaction(db.db, ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {
		return createMember(ctx, tx, todoListID, userID, MEMBER)
	})

}
