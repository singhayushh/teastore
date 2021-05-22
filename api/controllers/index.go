package controllers

import (
	"fmt"
	"teastore/api/auth"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// RenderHome ... render index.html on path '/'
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Teastore - Home",
	})
}

// RenderAbout ... render about.html on path '/about'
func RenderAbout(c *gin.Context) {
	c.HTML(200, "public_about.html", gin.H{
		"title": "Teastore - About Us",
	})
}

// RenderContact ... render contact.html on path '/contact'
func RenderContact(c *gin.Context) {
	c.HTML(200, "public_contact.html", gin.H{
		"title": "Teastore - Contact Us",
	})
}

// RenderRegister ...
func RenderRegister(c *gin.Context) {
	c.HTML(200, "public_register.html", gin.H{
		"title": "Register - Teastore",
	})
}

// RenderLogin ...
func RenderLogin(c *gin.Context) {
	c.HTML(200, "publuc_login.html", gin.H{
		"title": "Login - Teastore",
	})
}

// Register ... handler for POST /register
func (server *Server) Register(c *gin.Context) {
	user := models.User{}
	var err error

	fmt.Println(c)
	if err := c.ShouldBind(&user); err != nil {
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

	c.Redirect(301, "/login?message=success")
}

// Login ... handler for POST /login
func (server *Server) Login(c *gin.Context) {
	user := models.User{}
	var err error

	if err := c.ShouldBind(&user); err != nil {
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

// SignIn ... utility function used by Login
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
