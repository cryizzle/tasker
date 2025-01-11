package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPossibleNextStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   TodoStatus
		expected []TodoStatus
	}{
		{
			name:     "TODO status",
			status:   TODO,
			expected: []TodoStatus{ONGOING},
		},
		{
			name:     "ONGOING status",
			status:   ONGOING,
			expected: []TodoStatus{TODO, DONE},
		},
		{
			name:     "DONE status",
			status:   DONE,
			expected: []TodoStatus{ONGOING},
		},
		{
			name:     "Unknown status",
			status:   "UNKNOWN",
			expected: []TodoStatus{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo := Todo{Status: tt.status}
			actual := todo.GetPossibleNextStatus()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
func TestVerifyNextStatus(t *testing.T) {
	tests := []struct {
		name       string
		status     TodoStatus
		nextStatus TodoStatus
		expected   bool
	}{
		{
			name:       "Valid transition from TODO to ONGOING",
			status:     TODO,
			nextStatus: ONGOING,
			expected:   true,
		},
		{
			name:       "Invalid transition from TODO to DONE",
			status:     TODO,
			nextStatus: DONE,
			expected:   false,
		},
		{
			name:       "Valid transition from ONGOING to TODO",
			status:     ONGOING,
			nextStatus: TODO,
			expected:   true,
		},
		{
			name:       "Valid transition from ONGOING to DONE",
			status:     ONGOING,
			nextStatus: DONE,
			expected:   true,
		},
		{
			name:       "Invalid transition from DONE to TODO",
			status:     DONE,
			nextStatus: TODO,
			expected:   false,
		},
		{
			name:       "Valid transition from DONE to ONGOING",
			status:     DONE,
			nextStatus: ONGOING,
			expected:   true,
		},
		{
			name:       "Invalid transition from UNKNOWN to TODO",
			status:     "UNKNOWN",
			nextStatus: TODO,
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo := Todo{Status: tt.status}
			actual := todo.VerifyNextStatus(tt.nextStatus)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
