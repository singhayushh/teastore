package controllers

import (
	"fmt"
	"strconv"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// RenderHome ... render index.html on path '/'
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home | TEASTORE",
	})
}

// RenderAbout ... render about.html on path '/about'
func RenderAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "About | TEASTORE",
	})
}

// RenderContact ... render contact.html on path '/contact'
func RenderContact(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title": "Contact | TEASTORE",
	})
}

// RenderDashboard ... render dashboard.html on path '/dashboard'
func (server *Server) RenderDashboard(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	user := models.User{}
	fetchedUser, err := user.FetchByID(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	users, err := user.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	product := models.Product{}
	products, err := product.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	blog := models.Blog{}
	blogs, err := blog.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "dashboard.html", gin.H{
		"title":        "Dashboard | TEASTORE",
		"userCount":    len(*users),
		"productCount": len(*products),
		"blogCount":    len(*blogs),
		"productHits":  100,
		"blogHits":     100,
		"websiteHits":  100,
		"admin":        fetchedUser,
	})
}
