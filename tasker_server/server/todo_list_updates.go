package server

import (
	"io"

	"github.com/gin-gonic/gin"
)

func (srv Server) UpdateTodoList(c *gin.Context) {
	todoListID := c.Param("todo_list_id")

	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			if msg == todoListID {
				c.SSEvent("message", msg)
			}
			return true
		}
		return false
	})
}
