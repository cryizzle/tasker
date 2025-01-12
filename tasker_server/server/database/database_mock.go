package database

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetMembershipsForUser(ctx context.Context, userID uint64) ([]Member, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]Member), args.Error(1)
}

func (m *MockDatabase) GetTodoListMembers(ctx context.Context, todoListID uint64) ([]Member, error) {
	args := m.Called(ctx, todoListID)
	return args.Get(0).([]Member), args.Error(1)
}

func (m *MockDatabase) GetTodoEvents(ctx context.Context, todoID uint64) ([]TodoEvent, error) {
	args := m.Called(ctx, todoID)
	return args.Get(0).([]TodoEvent), args.Error(1)
}

func (m *MockDatabase) CreateTodoList(ctx context.Context, name string, token uuid.UUID, user *User) (uint64, error) {
	args := m.Called(ctx, name, token, user)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockDatabase) GetTodoList(ctx context.Context, query *TodoListQueryParam) (*TodoList, error) {
	args := m.Called(ctx, query)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*TodoList), args.Error(1)
}

func (m *MockDatabase) ListTodoLists(ctx context.Context, userID uint64) ([]TodoList, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]TodoList), args.Error(1)
}

func (m *MockDatabase) JoinTodoList(ctx context.Context, todoListID uint64, userID uint64) error {
	args := m.Called(ctx, todoListID, userID)
	return args.Error(0)
}

func (m *MockDatabase) GetTodos(ctx context.Context, todoListID uint64) ([]Todo, error) {
	args := m.Called(ctx, todoListID)
	return args.Get(0).([]Todo), args.Error(1)
}

func (m *MockDatabase) GetTodo(ctx context.Context, todoID uint64) (*Todo, error) {
	args := m.Called(ctx, todoID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Todo), args.Error(1)
}

func (m *MockDatabase) UpdateTodo(ctx context.Context, todo *Todo, user *User, nextStatus TodoStatus) error {
	args := m.Called(ctx, todo, user, nextStatus)
	return args.Error(0)
}

func (m *MockDatabase) CreateTodo(ctx context.Context, todo *Todo) (uint64, error) {
	args := m.Called(ctx, todo)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockDatabase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockDatabase) GetUserByID(ctx context.Context, id uint64) (*User, error) {
	args := m.Called(ctx, id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockDatabase) CreateUser(ctx context.Context, email string) (uint64, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(uint64), args.Error(1)
}
