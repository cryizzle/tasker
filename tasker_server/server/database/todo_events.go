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
	ID            uint64        `json:"id" db:"todo_event_id"`
	TodoID        uint64        `json:"todo_id" db:"todo_id"`
	OldValue      string        `json:"old_value" db:"old_value"`
	NewValue      string        `json:"new_value" db:"new_value"`
	EventType     TodoEventType `json:"event_type" db:"event_type"`
	CreatedAt     string        `json:"created_at" db:"created_at"`
	CreatedBy     uint64        `json:"-" db:"created_by"`
	CreatedByUser User          `json:"created_by"`
}

func (db Database) GetTodoEvents(ctx context.Context, todoID uint64) ([]TodoEvent, error) {
	var events []TodoEvent
	query := `SELECT * FROM todo_events WHERE todo_id = ? ORDER BY created_at DESC`
	err := db.db.SelectContext(ctx, &events, query, todoID)
	return events, err
}

func createTodoEvent(ctx context.Context, tx *sqlx.Tx, event *TodoEvent) error {
	query := `INSERT INTO todo_events (todo_id, old_value, new_value, event_type, created_by) VALUES (?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, event.TodoID, event.OldValue, event.NewValue, event.EventType, event.CreatedBy)
	return err
}
