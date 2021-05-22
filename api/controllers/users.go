package controllers

import (
	"fmt"
	"strconv"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// RenderUser ... display the user profile
func (server *Server) RenderUserByID(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
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
	c.HTML(200, "user_view.html", gin.H{
		"title": "Teastore - Account",
		"user":  fetchedUser,
	})
}

// RenderEditUser ... edit user page
func (server *Server) RenderEditUser(c *gin.Context) {
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
	c.HTML(200, "editUser.html", gin.H{
		"title": "Teastore - Edit Profile",
		"user":  fetchedUser,
	})
}

// UpdateUser updates the detials of the user sending the request
func (server *Server) UpdateUserByID(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{}

	if err := c.ShouldBind(&user); err != nil {
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
}

// DeleteUser removes the requesting user
func (server *Server) DeleteUserByID(c *gin.Context) {
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
}

// Logout ... Sets an invalid session id in the cookie instead of deleting it
func (server *Server) Logout(c *gin.Context) {
	c.SetCookie("session_id", "expired", 0, "/", "", false, false)
	c.Redirect(301, "/")
}
