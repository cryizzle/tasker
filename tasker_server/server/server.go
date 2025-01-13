package server

import (
	"errors"
	"log"
	"strconv"

	"github.com/cryizzle/tasker/tasker_server/server/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// It keeps a list of clients those are currently attached
// and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

type Server struct {
	Router *gin.Engine
	DB     database.DatabaseImpl
	Event  *Event

}

func NewEvent() (event *Event) {
	return &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}
}

func CreateServer(db *sqlx.DB) *Server {

	server := &Server{
		Router: gin.Default(),
		DB:     database.NewDatabase(db),
		Event:  NewEvent(),
	}
	return server
}

func (srv *Server) Routes(allowedOrigins []string) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowedOrigins
	corsConfig.AllowCredentials = true
	srv.Router.Use(cors.New(
		corsConfig,
	))

	userGroup := srv.Router.Group("/user")
	{
		userGroup.POST("/login", srv.UserLogin)
	}

	listGroup := srv.Router.Group("/list", AuthUser())
	{
		listGroup.POST("/create", srv.CreateTodoList)
		listGroup.POST("/join/:todo_list_token", srv.JoinTodoList)
		listGroup.GET("/all", srv.ListTodoLists)
		listGroup.GET("/:todo_list_id", srv.GetTodoList)
		// SSE
		listGroup.GET("/updates/:todo_list_id", SSEHeadersMiddleware(), srv.manageClientChannel(), srv.TodoListUpdates)
	}

	todoGroup := srv.Router.Group("/todo", AuthUser())
	{
		todoGroup.POST("/create", srv.CreateTodo)
		todoGroup.POST("/update/:todo_id", srv.UpdateTodo)
		todoGroup.GET("/events/:todo_id", srv.GetTodoEvents)
	}

}

func (server *Server) Start(port string) {
	// event listener in a separate go routine
	go server.listen()

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

func (srv *Server) listen() {
	for {
		select {
		// Add new available client
		case client := <-srv.Event.NewClients:
			srv.Event.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(srv.Event.TotalClients))

		// Remove closed client
		case client := <-srv.Event.ClosedClients:
			delete(srv.Event.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(srv.Event.TotalClients))

		// Broadcast message to client
		case eventMsg := <-srv.Event.Message:
			for clientMessageChan := range srv.Event.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func SSEHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

func (srv *Server) manageClientChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)

		// Send new connection to event server
		srv.Event.NewClients <- clientChan

		defer func() {
			// Drain client channel so that it does not block. Server may keep sending messages to this channel
			go func() {
				for range clientChan {
				}
			}()
			// Send closed connection to event server
			srv.Event.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func (srv *Server) SendTodoListUpdate(todoListID uint64) {
	srv.Event.Message <- strconv.FormatUint(todoListID, 10)
}
