package server

import (
	"log"
	"strconv"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
)

type UpdateTodoRequest struct {
	Status database.TodoStatus `json:"status"`
}

func (srv Server) UpdateTodo(c *gin.Context) {
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

	var request UpdateTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
	}

	user, err := srv.DB.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting user by ID"})
		return
	}

	todo, err := srv.DB.GetTodo(c.Request.Context(), todoID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error getting todo"})
		return
	}

	if !user.IsMember(todo.TodoListID) {
		c.JSON(403, gin.H{"error": "Unable to update todo - User is not a member of the todo list"})
		return
	}

	if !todo.VerifyNextStatus(request.Status) {
		c.JSON(400, gin.H{"error": "Invalid status transition"})
		return
	}

	err = srv.DB.UpdateTodo(c.Request.Context(), todo, user, request.Status)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error updating todo"})
		return
	}

	todo, err = srv.DB.GetTodo(c.Request.Context(), todoID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting todo"})
		return
	}

	c.JSON(200, todo)
}
