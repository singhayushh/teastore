package controllers

import (
	"github.com/gin-gonic/gin"
)

// Ping is a hello-world equivalent handler function for "/ping" route
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// RenderHome renders the home page for a get request on "/" route
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Gin is cool",
	})

}

// RenderRegister ...
func RenderRegister(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{
		"title": "Tea Store | Register",
	})
}

// RenderLogin ...
func RenderLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "Tea Store | Login",
	})
}
