package server

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	Email string `json:"email"`
}

func (srv Server) UserLogin(c *gin.Context) {
	var request UserLoginRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
	}

	user, err := srv.DB.GetUserByEmail(ctx, request.Email)

	if err == nil && user != nil {
		setActiveUserID(c, user.ID)
		c.JSON(200, gin.H{"user": user})
		return
	}

	if err != sql.ErrNoRows {
		c.JSON(500, gin.H{"error": "Error getting user by email"})
		return
	}

	log.Println("Error getting user by email: ", err)
	log.Println("User not found, creating new user")
	userID, err := srv.DB.CreateUser(ctx, request.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	user, err = srv.DB.GetUserByID(ctx, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting user by ID"})
		return
	}

	setActiveUserID(c, userID)
	c.JSON(200, gin.H{"user": user})

}

func setActiveUserID(c *gin.Context, userID uint64) {
	c.SetCookie(KEY_USER_ID, strconv.FormatUint(userID, 10), 0, "/", "localhost", false, true)
	log.Println("Cookie=", c.Request.Cookies())
}
