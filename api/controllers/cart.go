package controllers

import (
	"fmt"
	"strconv"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// RenderCart
func (server *Server) RenderCart(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	user := models.User{}
	_, err = user.FetchByID(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
}

// RenderCart
func (server *Server) RenderCheckout(c *gin.Context) {

}

// RenderCart
func (server *Server) AddtoCart(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	user := models.User{}
	_, err = user.FetchByID(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	type productInfo struct {
		ID   uint64 `gorm:"primary_key;auto_increment" json:"id" form:"id"`
		Path string `gorm:"unique" json:"path" form:"path"`
	}
	newProduct := new(productInfo)
	if err := c.ShouldBind(&newProduct); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	product := models.Product{}
	fetchedProduct, err := product.FetchByPath(server.DB, newProduct.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	fmt.Println(fetchedProduct)

	updatedUser, err := user.AddtoCart(server.DB, uid, fetchedProduct)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": updatedUser})
}

// RenderCart
func (server *Server) RemovefromCart(c *gin.Context) {

}

// RenderCart
func (server *Server) Checkout(c *gin.Context) {

}
