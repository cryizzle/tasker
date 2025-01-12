package server

import (
	"bytes"
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

func TestCreateTodoList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockTodoList := database.TodoList{
		ID:   1,
		Name: "Test Todo List",
	}

	tests := []struct {
		name               string
		requestBody        gin.H
		setupMocks         func(m *database.MockDatabase)
		expectedStatusCode int
		expectedResponse   gin.H
	}{
		{
			name: "Missing name",
			requestBody: gin.H{
				"name": "",
			},
			setupMocks:         func(m *database.MockDatabase) {},
			expectedStatusCode: 400,
			expectedResponse:   gin.H{"error": "Name is required"},
		},
		{
			name: "Error getting user by ID",
			requestBody: gin.H{
				"name": "Test Todo List",
			},
			setupMocks: func(m *database.MockDatabase) {
				m.On("GetUserByID", mock.Anything, mock.Anything).Return(nil, errors.New("user not found"))
			},
			expectedStatusCode: 500,
			expectedResponse:   gin.H{"error": "Error getting user by ID"},
		},
		{
			name: "Error creating todo list",
			requestBody: gin.H{
				"name": "Test Todo List",
			},
			setupMocks: func(m *database.MockDatabase) {
				m.On("GetUserByID", mock.Anything, mock.Anything).Return(&database.User{}, nil)
				m.On("CreateTodoList", mock.Anything, "Test Todo List", mock.Anything, mock.Anything).Return(uint64(0), errors.New("creation error"))
			},
			expectedStatusCode: 500,
			expectedResponse:   gin.H{"error": "Error creating todo list"},
		},
		{
			name: "Error getting todo list",
			requestBody: gin.H{
				"name": "Test Todo List",
			},
			setupMocks: func(m *database.MockDatabase) {
				m.On("GetUserByID", mock.Anything, mock.Anything).Return(&database.User{}, nil)
				m.On("CreateTodoList", mock.Anything, "Test Todo List", mock.Anything, mock.Anything).Return(uint64(1), nil)
				m.On("GetTodoList", mock.Anything, mock.Anything).Return(nil, errors.New("get todo list error"))
			},
			expectedStatusCode: 500,
			expectedResponse:   gin.H{"error": "Error getting todo list"},
		},
		{
			name: "Successful creation",
			requestBody: gin.H{
				"name": "Test Todo List",
			},
			setupMocks: func(m *database.MockDatabase) {
				m.On("GetUserByID", mock.Anything, mock.Anything).Return(&database.User{}, nil)
				m.On("CreateTodoList", mock.Anything, "Test Todo List", mock.Anything, mock.Anything).Return(uint64(1), nil)
				m.On("GetTodoList", mock.Anything, mock.Anything).Return(&mockTodoList, nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   gin.H{"todo_list": mockTodoList},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := new(database.MockDatabase)
			tt.setupMocks(mockDB)

			srv := Server{DB: mockDB}

			router := gin.Default()
			router.POST("/list/create",
				func(c *gin.Context) {
					c.Set(KEY_USER_ID, "1")
					c.Next()
				}, srv.CreateTodoList)

			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/list/create", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			expectedBodyJSON, err := json.Marshal(tt.expectedResponse)
			assert.NoError(t, err)
			assert.JSONEq(t, string(expectedBodyJSON), w.Body.String())
		})
	}
}
