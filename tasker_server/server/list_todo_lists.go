package server

import "github.com/gin-gonic/gin"

type ListTodoListsRequest struct {
	UserIDRequest
}

func (srv Server) ListTodoLists(c *gin.Context) {
	var request ListTodoListsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	todoLists, err := srv.DB.ListTodoLists(c.Request.Context(), request.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error listing todo lists"})
		return
	}

	c.JSON(200, gin.H{"todo_lists": todoLists})
}
