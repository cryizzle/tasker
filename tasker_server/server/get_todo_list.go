package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetTodoListRequest struct {
	UserIDRequest
}

func (srv Server) GetTodoList(c *gin.Context) {
	param := c.Param("todo_list_id")
	todoListID, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo list ID"})
		return
	}

	var request GetTodoListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	todoList, err := srv.DB.GetTodoList(c.Request.Context(), todoListID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": todoList})
}
