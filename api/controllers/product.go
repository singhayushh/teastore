package controllers

import (
	"fmt"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// AddProduct ... adds new product in the db
func (server *Server) AddProduct(c *gin.Context) {
	product := models.Product{}
	var err error

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Check if all parameters have been inputted
	err = product.Validate("")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// _ to be changed and used for page rendering
	_, err = product.Save(server.DB)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// ShowProduct fetches data of the product by id
func (server *Server) ShowProduct(c *gin.Context) {
	path := c.Param("path")
	product := fmt.Sprint("tea-", path, ".html")
	// product := models.Product{}
	// fetchedProduct, err := product.FetchByID(server.DB, path)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err})
	// 	return
	// }
	// c.JSON(200, gin.H{"product": fetchedProduct})
	c.HTML(200, product, gin.H{
		"title": "Gin is cool",
	})
}

// UpdateProduct updates the detials of the product
func (server *Server) UpdateProduct(c *gin.Context) {
	product := models.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err := product.Update(server.DB, product.Path)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": product})
	return

}

// DeleteProduct removes the requested product
func (server *Server) DeleteProduct(c *gin.Context) {
	product := models.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	_, err := product.Delete(server.DB, product.Path)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"updated": "success"})
	return
}
