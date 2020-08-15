package controllers

import (
	"fmt"
	"strconv"
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
		fmt.Println(user)
		return
	}

	// Check if all parameters have been inputted
	err = user.Validate("")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Password Hashing
	err = user.EncryptPassword()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// _ to be changed and used for page rendering
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
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	err = user.Validate("login")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	sessionID, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		c.JSON(206, gin.H{"message": "failed"})
	} else {
		c.SetCookie("session_id", sessionID, 7.884e+6, "/", "", false, false)
		c.Redirect(301, "/")
	}
}

// SignIn is a utility function used by Login() which finds user via mail, checks his password and calls auth.CreateSession()
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

	return auth.CreateSession(user.ID, user.Type)
}

// ShowUser fetches data of the user by his id
func (server *Server) ShowUser(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	fmt.Println("uid", uid)
	user := models.User{}
	fetchedUser, err := user.FetchByID(server.DB, uid)
	c.JSON(200, gin.H{"user": fetchedUser})
}

// UpdateUser updates the detials of the user sending the request
func (server *Server) UpdateUser(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err = user.Update(server.DB, uid)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": user})
	return

}

// DeleteUser removes the requesting user
func (server *Server) DeleteUser(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{}

	_, err = user.Delete(server.DB, uid)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
	return

}
