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
		"title": "Home | TEASTORE",
	})

}

func RenderAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "About | TEASTORE",
	})

}

func RenderContact(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title": "Contact | TEASTORE",
	})

}
