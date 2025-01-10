package database

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type TodoList struct {
	ID      uint64    `json:"id" db:"todo_list_id"`
	Token   uuid.UUID `json:"token" db:"token"`
	Name    string    `json:"name" db:"name"`
	Members []Member  `json:"members"`
	Todos   []Todo    `json:"todos"`
}

func (db Database) CreateTodoList(ctx context.Context, name string, token uuid.UUID, user *User) (uint64, error) {

	return WithTransactionRet(db.db, ctx, nil, func(ctx context.Context, tx *sqlx.Tx) (uint64, error) {

		log.Println(token)

		query := `INSERT INTO todo_lists (name, token) VALUES (?, ?)`
		result, err := tx.ExecContext(ctx, query, name, token.Bytes())

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

type TodoListQueryParam struct {
	TodoListID    uint64
	TodoListToken uuid.UUID
}

func (db Database) GetTodoList(ctx context.Context, query *TodoListQueryParam) (*TodoList, error) {
	var todoList TodoList
	var queryParam []string
	var args []any

	if query == nil {
		return nil, errors.New("no query parameters provided")
	}

	if query.TodoListID != 0 {
		queryParam = append(queryParam, "todo_list_id = ?")
		args = append(args, query.TodoListID)
	}

	if query.TodoListToken != uuid.Nil {
		queryParam = append(queryParam, "token = ?")
		args = append(args, query.TodoListToken.Bytes())
	}

	if len(queryParam) == 0 || len(args) == 0 {
		return nil, errors.New("no query parameters provided")
	}

	queryString := `SELECT * FROM todo_lists WHERE ` + strings.Join(queryParam, " AND ")
	err := db.db.GetContext(ctx, &todoList, queryString, args...)

	if err != nil {
		return nil, err
	}

	members, err := db.GetTodoListMembers(ctx, todoList.ID)

	if err != nil {
		return nil, err
	}
	todoList.Members = members

	todos, err := db.GetTodos(ctx, todoList.ID)
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
	tdl.name as name,
	tdl.token as token
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
