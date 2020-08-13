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
	c.HTML(200, "temp.html", gin.H{
		"title": "Gin is cool",
	})
}

// RenderProfile renders the profile page for a get request on "/profile" route
func RenderProfile(c *gin.Context) {
	c.HTML(200, "temp2.html", gin.H{
		"title": "Logged in?",
	})
}
