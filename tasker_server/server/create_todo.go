package server

import (
	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Description string `json:"description"`
	TodoListID  uint64 `json:"todo_list_id"`
}

func (srv Server) CreateTodo(c *gin.Context) {
	var request CreateTodoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if request.Description == "" {
		c.JSON(400, gin.H{"error": "Description is required"})
		return
	}

	userID, err := GetAuthenticatedUser(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	user, err := srv.DB.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting user by ID"})
		return
	}

	if !user.IsMember(request.TodoListID) {
		c.JSON(403, gin.H{"error": "Unable to create todo - User is not a member of the todo list"})
		return
	}

	todoID, err := srv.DB.CreateTodo(c.Request.Context(), &database.Todo{
		Description: request.Description,
		TodoListID:  request.TodoListID,
		Status:      database.TODO,
		CreatedBy:   user.ID,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	todo, err := srv.DB.GetTodo(c.Request.Context(), todoID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo"})
		return
	}

	srv.SendTodoListUpdate(todo.TodoListID)

	c.JSON(200, gin.H{"todo": todo})
}
