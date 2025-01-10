package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (srv Server) ListTodoLists(c *gin.Context) {

	userID, err := GetAuthenticatedUser(c)
	if err != nil {
		c.JSON(400, err)
		return
	}

	todoLists, err := srv.DB.ListTodoLists(c.Request.Context(), userID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error listing todo lists"})
		return
	}

	c.JSON(200, gin.H{"todo_lists": todoLists})
}
