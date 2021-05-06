package controllers

import (
	"fmt"
	"teastore/api/models"

	"github.com/gin-gonic/gin"
)

// ShowAllProducts ...
func (server *Server) RenderAllProducts(c *gin.Context) {
	product := models.Product{}
	products, err := product.FetchAll(server.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.HTML(200, "listProduct.html", gin.H{
		"title":    "Dashboard | TEASTORE",
		"products": products,
	})
}

// AddProduct ... adds new product in the db
func (server *Server) AddProduct(c *gin.Context) {
	product := models.Product{}
	var err error

	if err := c.ShouldBind(&product); err != nil {
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

// RenderProduct fetches data of the product by id (path)
func (server *Server) RenderProduct(c *gin.Context) {
	id := c.Param("id")
	product := models.Product{}
	fetchedProduct, err := product.FetchByID(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "viewProduct.html", gin.H{
		"title":   "View Product | Teastore",
		"product": fetchedProduct,
	})
}

// RenderEditProduct fetches data of the product by id (path)
func (server *Server) RenderEditProduct(c *gin.Context) {
	id := c.Param("id")
	product := models.Product{}
	fetchedProduct, err := product.FetchByID(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.HTML(200, "editProduct.html", gin.H{
		"title":   "Edit Product | Teastore",
		"product": fetchedProduct,
	})
}

// UpdateProductByID updates the detials of the product
func (server *Server) UpdateProductByID(c *gin.Context) {
	id := c.Param("id")
	product := models.Product{}
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	_, err := product.Update(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"updated": product})
}

// DeleteProductByID removes the requested product
func (server *Server) DeleteProductByID(c *gin.Context) {
	id := c.Param("id")
	product := models.Product{}
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	_, err := product.Delete(server.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"updated": "success"})
}
