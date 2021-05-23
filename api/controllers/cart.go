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
	cart := models.Cart{}
	fetchedCart, err := cart.FetchCart(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	fmt.Println(fetchedCart)
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

	type cartItem struct {
		UserID    uint64 `gorm:"" json:"uid" form:"uid"`
		ProductID uint64 `gorm:"" json:"pid" form:"pid"`
		Quantity  int    `gorm:"" json:"qty" form:"qty"`
	}
	newItem := new(cartItem)
	if err := c.ShouldBind(&newItem); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	newItem.UserID = uid

	fmt.Println(newItem)

	cart := models.Cart{}
	err = cart.AddtoCart(server.DB, newItem.UserID, newItem.ProductID, newItem.Quantity)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	getCart, err := cart.FetchCart(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": getCart})
}

// RenderCart
func (server *Server) RemovefromCart(c *gin.Context) {

}

// RenderCart
func (server *Server) Checkout(c *gin.Context) {

}
