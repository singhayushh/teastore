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

	c.HTML(200, "user_cart.html", gin.H{
		"title": "Teastore - My Cart",
		"cart":  fetchedCart,
	})
}

// RenderCart
func (server *Server) RenderCheckout(c *gin.Context) {

}

// RenderCart
func (server *Server) AddtoCart(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid, err := strconv.ParseUint(fmt.Sprintf("%v", uidInterface), 10, 64)

	if err != nil {
		fmt.Println(err)
		c.Redirect(301, "/login")
	}
	user := models.User{}
	_, err = user.FetchByID(server.DB, uid)
	if err != nil {
		fmt.Println(err)
		c.Redirect(301, "/login")
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

	cart := models.Cart{}
	err = cart.AddtoCart(server.DB, newItem.UserID, newItem.ProductID, newItem.Quantity)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	_, err = cart.FetchCart(server.DB, uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.Redirect(301, "/cart")
}

// RenderCart
func (server *Server) RemovefromCart(c *gin.Context) {

}

// RenderCart
func (server *Server) Checkout(c *gin.Context) {

}
