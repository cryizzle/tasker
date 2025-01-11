package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TodoEventType string

const (
	TODO_CREATED        TodoEventType = "TODO_CREATED"
	TODO_STATUS_CHANGED TodoEventType = "TODO_STATUS_CHANGED"
)

type TodoEvent struct {
	ID        uint64        `json:"id" db:"todo_event_id"`
	TodoID    uint64        `json:"todo_id" db:"todo_id"`
	OldValue  string        `json:"old_value" db:"old_value"`
	NewValue  string        `json:"new_value" db:"new_value"`
	EventType TodoEventType `json:"event_type" db:"event_type"`
	CreatedAt string        `json:"created_at" db:"created_at"`
	CreatedBy uint64        `json:"created_by" db:"created_by"`
	User
}

func (db Database) GetTodoEvents(ctx context.Context, todoID uint64) ([]TodoEvent, error) {
	var events []TodoEvent
	query := `SELECT
	todo_events.todo_event_id as todo_event_id,
	todo_events.todo_id as todo_id,
	todo_events.old_value as old_value,
	todo_events.new_value as new_value,
	todo_events.event_type as event_type,
	todo_events.created_at as created_at,
	todo_events.created_by as created_by,
	users.user_id as user_id,
	users.email as email
	FROM todo_events
	JOIN users ON todo_events.created_by = users.user_id
	WHERE todo_id = ? ORDER BY created_at DESC`
	err := db.db.SelectContext(ctx, &events, query, todoID)
	return events, err
}

func createTodoEvent(ctx context.Context, tx *sqlx.Tx, event *TodoEvent) error {
	query := `INSERT INTO todo_events (todo_id, old_value, new_value, event_type, created_by) VALUES (?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, event.TodoID, event.OldValue, event.NewValue, event.EventType, event.CreatedBy)
	return err
}
