package server

import (
	"log"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CreateTodoListRequest struct {
	Name string `json:"name"`
}

func (srv Server) CreateTodoList(c *gin.Context) {
	var request CreateTodoListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
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

	todoListToken, err := uuid.NewV4()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating todo list token"})
		return
	}

	todoListID, err := srv.DB.CreateTodoList(c.Request.Context(), request.Name, todoListToken, user)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error creating todo list"})
		return
	}

	todoList, err := srv.DB.GetTodoList(c.Request.Context(),
		&database.TodoListQueryParam{
			TodoListID:    todoListID,
			TodoListToken: todoListToken,
		})

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error getting todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": todoList})

}
