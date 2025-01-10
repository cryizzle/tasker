package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (srv Server) GetTodoEvents(c *gin.Context) {

	param := c.Param("todo_id")
	todoID, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
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

	todo, err := srv.DB.GetTodo(c.Request.Context(), todoID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo"})
		return
	}

	if !user.IsMember(todo.TodoListID) {
		c.JSON(403, gin.H{"error": "Unable to get todo events - User is not a member of the todo list"})
		return
	}

	todoEvents, err := srv.DB.GetTodoEvents(c.Request.Context(), todo.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo events"})
		return
	}

	c.JSON(200, gin.H{"todo_events": todoEvents})
}
