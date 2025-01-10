package server

import (
	"strconv"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
)

func (srv Server) GetTodoList(c *gin.Context) {
	param := c.Param("todo_list_id")
	todoListID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo list ID"})
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

	if !user.IsMember(todoListID) {
		c.JSON(403, gin.H{"error": "Unable to view todo - User is not a member of the todo list"})
		return
	}

	todoList, err := srv.DB.GetTodoList(c.Request.Context(), &database.TodoListQueryParam{
		TodoListID: todoListID,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": todoList})
}
