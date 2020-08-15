package controllers

import (
	"fmt"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// CreateBlog ... adds new blog in the db
func (server *Server) CreateBlog(c *gin.Context) {
	blog := models.Blog{}
	var err error

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Check if all parameters have been inputted
	err = blog.Validate("")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// _ to be changed and used for page rendering
	_, err = blog.Save(server.DB)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// ReadBlog fetches data of the blog by id
func (server *Server) ReadBlog(c *gin.Context) {
	path := c.Param("path")
	blog := models.Blog{}
	fetchedBlog, err := blog.FetchByID(server.DB, path)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"blog": fetchedBlog})
}

// UpdateBlog updates the detials of the blog
func (server *Server) UpdateBlog(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err := blog.Update(server.DB, blog.Path)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": blog})
	return

}

// DeleteBlog removes the requested blog
func (server *Server) DeleteBlog(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err := blog.Delete(server.DB, blog.Path)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": "success"})
	return
}
