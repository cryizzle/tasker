package server

import (
	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
)

type CreateTodoListRequest struct {
	Name string `json:"name"`
	UserIDRequest
}

func (srv Server) CreateTodoList(c *gin.Context) {
	var request CreateTodoListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := srv.DB.GetUserByID(c.Request.Context(), request.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting user by ID"})
		return
	}

	todoListID, err := srv.DB.CreateTodoList(c.Request.Context(), request.Name, user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": database.TodoList{ID: todoListID}})

}
