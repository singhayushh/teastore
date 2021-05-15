package controllers

import (
	"fmt"
	"html/template"
	"teastore/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// RenderAllBlogs ...
func (server *Server) RenderAllBlogs(c *gin.Context) {
	blog := models.Blog{}
	blogs, err := blog.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.HTML(200, "listBlog.html", gin.H{
		"title":         "Blogs | TEASTORE",
		"blogs":         blogs,
		"loadDatatable": true,
	})
}

// CreateBlog ... adds new blog in the db
func (server *Server) CreateBlog(c *gin.Context) {
	blog := models.Blog{}
	var err error

	if err := c.ShouldBind(&blog); err != nil {
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

// RenderBlog fetches data of the blog by id
func (server *Server) RenderBlog(c *gin.Context) {
	id := c.Param("id")
	blog := models.Blog{}
	fetchedBlog, err := blog.FetchByID(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "viewBlog.html", gin.H{
		"title":       "Blog | TEASTORE",
		"blogTitle":   template.HTML(fetchedBlog.Title),
		"blogContent": template.HTML(fetchedBlog.Text),
		"blogCover":   fetchedBlog.Cover,
		"blogAuthor":  fetchedBlog.Author,
	})
}

// RenderAddBlog
func (server *Server) RenderAddBlog(c *gin.Context) {
	c.HTML(200, "addBlog.html", gin.H{
		"title":      "Add blog | Teastore",
		"loadEditor": true,
	})
}

// RenderEditBlog fetches data of the blog by id (path)
func (server *Server) RenderEditBlog(c *gin.Context) {
	id := c.Param("id")
	blog := models.Blog{}
	fetchedBlog, err := blog.FetchByID(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "editBlog.html", gin.H{
		"title":      "Edit blog | Teastore",
		"blog":       fetchedBlog,
		"loadEditor": true,
	})
}

// UpdateBlogByID updates the detials of the blog
func (server *Server) UpdateBlogByID(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	currentTime := time.Now()
	blog.UpdatedAt = currentTime.Format("2006-01-02")

	_, err := blog.Update(server.DB, blog.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": blog})
}

// DeleteBlogByID removes the requested blog
func (server *Server) DeleteBlogByID(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err := blog.Delete(server.DB, blog.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": "success"})
}
