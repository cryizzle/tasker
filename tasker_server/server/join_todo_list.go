package server

import (
	"strconv"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
)

type JoinTodoListRequest struct {
	UserIDRequest
}

func (srv Server) JoinTodoList(c *gin.Context) {

	var request JoinTodoListRequest
	pathParam := c.Param("todo_list_id")
	todoListID, err := strconv.ParseUint(pathParam, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo list ID"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := srv.DB.GetUserByID(c.Request.Context(), request.UserID)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting user by ID"})
		return
	}

	if user.IsMember(todoListID) {
		c.JSON(200, gin.H{"message": "User is already a member of this todo list"})
		return
	}

	err = srv.DB.JoinTodoList(c.Request.Context(), todoListID, user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error joining todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": database.TodoList{ID: todoListID}})

}
