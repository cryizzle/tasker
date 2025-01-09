package database

import (
	"context"
	"slices"
	"time"

	"github.com/jmoiron/sqlx"
)

type TodoStatus string

const (
	TODO    TodoStatus = "TODO"
	ONGOING TodoStatus = "ONGOING"
	DONE    TodoStatus = "DONE"
)

type Todo struct {
	ID            uint64     `json:"id" db:"todo_id"`
	TodoListID    uint64     `json:"todo_list_id" db:"todo_list_id"`
	Description   string     `json:"description" db:"description"`
	Status        TodoStatus `json:"status" db:"status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy     uint64     `json:"-" db:"created_by"`
	CreatedByUser *User      `json:"created_by"`
}

func (t Todo) VerifyNextStatus(nextStatus TodoStatus) bool {
	possibleNextStatus := t.GetPossibleNextStatus()
	return slices.Contains(possibleNextStatus, nextStatus)
}

func (t Todo) GetPossibleNextStatus() []TodoStatus {
	switch t.Status {
	case TODO:
		return []TodoStatus{ONGOING}
	case ONGOING:
		return []TodoStatus{TODO, DONE}
	case DONE:
		return []TodoStatus{ONGOING}
	default:
		return []TodoStatus{}
	}
}

func (db Database) GetTodos(ctx context.Context, todoListID uint64) ([]Todo, error) {
	var todos []Todo
	query := `SELECT * FROM todos WHERE todo_list_id = ?`

	err := db.db.SelectContext(ctx, &todos, query, todoListID)
	return todos, err

}

func (db Database) GetTodo(ctx context.Context, todoID uint64) (*Todo, error) {
	var todo Todo
	query := `SELECT * FROM todos WHERE todo_id = ?`

	err := db.db.GetContext(ctx, &todo, query, todoID)
	return &todo, err
}

func (db Database) UpdateTodo(ctx context.Context, todo *Todo, user *User, nextStatus TodoStatus) error {

	return WithTransaction(db.db, ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {

		var _temp Todo
		err := tx.GetContext(ctx, &_temp, `SELECT * FROM todos WHERE todo_id = ? FOR UPDATE`, todo.ID)
		if err != nil {
			return err
		}

		query := `UPDATE todos SET status = ? WHERE todo_id = ?`
		_, err = tx.ExecContext(ctx, query, nextStatus, todo.ID)

		if err != nil {
			return err
		}

		return createTodoEvent(ctx, tx, &TodoEvent{
			TodoID:    todo.ID,
			OldValue:  string(todo.Status),
			NewValue:  string(nextStatus),
			EventType: TODO_STATUS_CHANGED,
			CreatedBy: user.ID,
		})

	})

}

func (db Database) CreateTodo(ctx context.Context, todo *Todo) (uint64, error) {

	return WithTransactionRet(db.db, ctx, nil, func(ctx context.Context, tx *sqlx.Tx) (uint64, error) {
		query := `INSERT INTO todos (todo_list_id, description, status, created_by) VALUES (?, ?, ?, ?)`
		result, err := db.db.ExecContext(ctx, query, todo.TodoListID, todo.Description, todo.Status, todo.CreatedBy)

		if err != nil {
			return 0, err
		}

		todoID, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}

		err = createTodoEvent(ctx, tx, &TodoEvent{
			TodoID:    uint64(todoID),
			OldValue:  "",
			NewValue:  todo.Description,
			EventType: TODO_CREATED,
			CreatedBy: todo.CreatedBy,
		})

		if err != nil {
			return 0, err
		}

		return uint64(todoID), nil
	})

}
