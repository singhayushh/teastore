package controllers

import (
	"fmt"
	"html/template"
	"strconv"
	"teastore/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// RenderBlogs ...
func (server *Server) RenderBlogs(c *gin.Context) {
	blog := models.Blog{}
	blogs, err := blog.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.HTML(200, "blog_feed.html", gin.H{
		"title": "Teastore - Blogs",
		"blogs": blogs,
	})
}

// RenderBlog fetches data of the blog by path
func (server *Server) RenderBlogByPath(c *gin.Context) {
	path := c.Param("path")
	blog := models.Blog{}
	fetchedBlog, err := blog.FetchByPath(server.DB, path)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "blog_view.html", gin.H{
		"title":       "Blog | TEASTORE",
		"blogTitle":   template.HTML(fetchedBlog.Title),
		"blogContent": template.HTML(fetchedBlog.Text),
		"blogCover":   fetchedBlog.Cover,
		"blogAuthor":  fetchedBlog.Author,
	})
}

// RenderAddBlog
func (server *Server) RenderAddBlog(c *gin.Context) {
	c.HTML(200, "blog_add.html", gin.H{
		"title":      "Add blog | Teastore",
		"loadEditor": true,
	})
}

// RenderEditBlog fetches data of the blog by id (path)
func (server *Server) RenderEditBlog(c *gin.Context) {
	path := c.Param("path")
	blog := models.Blog{}
	fetchedBlog, err := blog.FetchByPath(server.DB, path)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "blog_edit.html", gin.H{
		"title":      "Edit blog | Teastore",
		"blog":       fetchedBlog,
		"loadEditor": true,
	})
}

// CreateBlog ... adds new blog in the db
func (server *Server) CreateBlog(c *gin.Context) {
	blog := models.Blog{}
	var err error

	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	// Check if all parameters have been inputted
	err = blog.Validate("")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	// _ to be changed and used for page rendering
	_, err = blog.Save(server.DB)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.Redirect(301, "/blogs/view/"+blog.Path)
}

// UpdateBlogByID updates the detials of the blog
func (server *Server) UpdateBlogByID(c *gin.Context) {
	uidInterface := c.Param("id")
	id, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	blog := models.Blog{}

	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	currentTime := time.Now()
	blog.UpdatedAt = currentTime.Format("2006-01-02")

	_, err = blog.Update(server.DB, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.Redirect(301, "/blogs/view/"+blog.Path)
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

	c.Redirect(301, "/dashboard/blogs")
}
