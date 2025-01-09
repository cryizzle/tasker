package server

import (
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
	srv.Router.Use(cors.New(
		corsConfig,
	))

	userGroup := srv.Router.Group("/user")
	{
		userGroup.POST("/login", srv.UserLogin)
	}

	listGroup := srv.Router.Group("/list")
	{
		listGroup.POST("/create", srv.CreateTodoList)
		listGroup.POST("/join/:todo_list_id", srv.JoinTodoList)
		listGroup.GET("/all", srv.ListTodoLists)
		listGroup.GET("/:todo_list_id", srv.GetTodoList)
	}

	todoGroup := srv.Router.Group("/todo")
	{
		todoGroup.POST("/create", srv.CreateTodo)
		todoGroup.POST("/update/:todo_id", srv.UpdateTodo)
		todoGroup.GET("/events/:todo_id", srv.GetTodoEvents)
	}

}

func (server *Server) Start(port string) {
	server.Router.Run(port)
}
