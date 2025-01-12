package server

import (
	"log"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (srv Server) JoinTodoList(c *gin.Context) {

	pathParam := c.Param("todo_list_token")
	todoListToken := uuid.FromStringOrNil(pathParam)

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

	todoList, err := srv.DB.GetTodoList(c.Request.Context(), &database.TodoListQueryParam{
		TodoListToken: todoListToken,
	})
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error getting todo list by token"})
		return
	}

	if user.IsMember(todoList.ID) {
		c.JSON(500, gin.H{"error": "User is already a member of this todo list"})
		return
	}

	err = srv.DB.JoinTodoList(c.Request.Context(), todoList.ID, user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error joining todo list"})
		return
	}

	c.JSON(200, gin.H{"todo_list": todoList})

}
