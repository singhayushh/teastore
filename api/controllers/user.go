package controllers

import (
	"teastore/api/auth"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// Register is the handler for the register route
func (server *Server) Register(c *gin.Context) {
	user := models.User{}
	var err error

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	err = user.Validate("")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err = user.Save(server.DB)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// Login is the handler for the login route
func (server *Server) Login(c *gin.Context) {
	user := models.User{}
	var err error

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error1": err})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error2": err})
		return
	}

	sessionID, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		c.JSON(200, gin.H{"message": "failed"})
	} else {
		c.SetCookie("session_id", sessionID, 7.884e+6, "/", "", false, false)
		c.JSON(200, gin.H{"message": "success"})
	}
}

// SignIn finds user via mail, checks his password and calls auth.CreateSession()
func (server *Server) SignIn(email, password string) (string, error) {
	var err error

	var user = models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil {
		return "", err
	}

	return auth.CreateSession(user.Email)
}
