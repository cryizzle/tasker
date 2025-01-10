package server

import (
	"errors"
	"strconv"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Router *gin.Engine
	DB     *database.Database
}

func CreateServer(db *sqlx.DB) *Server {

	server := &Server{
		Router: gin.Default(),
		DB:     database.NewDatabase(db),
	}
	return server
}

func (srv *Server) Routes() {

	corsConfig := cors.DefaultConfig()
	// TODO: move to env
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowCredentials = true
	srv.Router.Use(cors.New(
		corsConfig,
	))

	userGroup := srv.Router.Group("/user")
	{
		userGroup.POST("/login", srv.UserLogin)
	}

	listGroup := srv.Router.Group("/list")
	{
		listGroup.POST("/create", AuthUser(), srv.CreateTodoList)
		listGroup.POST("/join/:todo_list_token", AuthUser(), srv.JoinTodoList)
		listGroup.GET("/all", AuthUser(), srv.ListTodoLists)
		listGroup.GET("/:todo_list_id", AuthUser(), srv.GetTodoList)
	}

	todoGroup := srv.Router.Group("/todo")
	{
		todoGroup.POST("/create", AuthUser(), srv.CreateTodo)
		todoGroup.POST("/update/:todo_id", AuthUser(), srv.UpdateTodo)
		todoGroup.GET("/events/:todo_id", AuthUser(), srv.GetTodoEvents)
	}

}

func (server *Server) Start(port string) {
	server.Router.Run(port)
}

const KEY_USER_ID = "user_id"

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := c.Cookie(KEY_USER_ID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set(KEY_USER_ID, userID)
		c.Next()
	}
}

func GetAuthenticatedUser(c *gin.Context) (uint64, error) {
	userIDString, exists := c.Get(KEY_USER_ID)
	if !exists {
		return 0, errors.New("user is not logged in")
	}
	userID, err := strconv.ParseUint(userIDString.(string), 10, 64)
	if err != nil {
		return 0, errors.New("invalid user ID")
	}

	return userID, nil
}
