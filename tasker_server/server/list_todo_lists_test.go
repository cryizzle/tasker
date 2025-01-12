package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListTodoLists(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockTodoList := database.TodoList{
		ID:   1,
		Name: "Test Todo List",
	}

	tests := []struct {
		name           string
		userID         string
		mockDBResponse []database.TodoList
		mockDBError    error
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			name:           "successful list",
			userID:         "1",
			mockDBResponse: []database.TodoList{mockTodoList},
			mockDBError:    nil,
			expectedStatus: http.StatusOK,
			expectedBody:   gin.H{"todo_lists": []database.TodoList{mockTodoList}},
		},
		{
			name:           "database error",
			userID:         "1",
			mockDBResponse: nil,
			mockDBError:    errors.New("db error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"error": "Error listing todo lists"},
		},
		{
			name:           "unauthenticated user",
			userID:         "",
			mockDBResponse: nil,
			mockDBError:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(database.MockDatabase)
			srv := Server{DB: mockDB}

			if tt.userID != "" {
				mockDB.On("ListTodoLists", mock.Anything, mock.Anything).Return(tt.mockDBResponse, tt.mockDBError)
			}

			r := gin.Default()
			r.GET("/todo_lists", func(c *gin.Context) {
				if tt.userID != "" {
					c.Set(KEY_USER_ID, tt.userID)
				}
				c.Next()
			},
				srv.ListTodoLists,
			)

			req, _ := http.NewRequest(http.MethodGet, "/todo_lists", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			expectedBodyJSON, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)
			assert.JSONEq(t, string(expectedBodyJSON), w.Body.String())
		})
	}
}
