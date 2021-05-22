package controllers

import (
	"fmt"
	"strconv"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

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
	c.HTML(200, "admin-dashboard.html", gin.H{
		"title":        "Teastore - Dashboard",
		"userCount":    len(*users),
		"productCount": len(*products),
		"blogCount":    len(*blogs),
		"productHits":  100,
		"blogHits":     100,
		"websiteHits":  100,
		"admin":        fetchedUser,
	})
}

func (server *Server) RenderUserDashboard(c *gin.Context) {
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
	c.HTML(200, "user_admin.html", gin.H{
		"title":         "Teastore - Users",
		"users":         users,
		"loadDatatable": true,
		"admin":         fetchedUser,
	})
}

func (server *Server) RenderBlogDashboard(c *gin.Context) {
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
	blog := models.Blog{}
	blogs, err := blog.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.HTML(200, "blog_admin.html", gin.H{
		"title":         "Teastore - Blogs",
		"blogs":         blogs,
		"loadDatatable": true,
		"admin":         fetchedUser,
	})
}

func (server *Server) RenderProductDashboard(c *gin.Context) {
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
	product := models.Product{}
	products, err := product.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.HTML(200, "product_admin.html", gin.H{
		"title":         "Teastore - Products",
		"products":      products,
		"loadDatatable": true,
		"admin":         fetchedUser,
	})
}
