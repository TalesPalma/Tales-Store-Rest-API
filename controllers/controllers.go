package controllers

import (
	"net/http"

	"github.com/TalesPalma/gin-golang-rest/models"
	"github.com/TalesPalma/gin-golang-rest/services"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, services.FindProducts())
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(
		200,
		services.FindProductsById(id),
	)
}

func PostProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !newProduct.ValidateFieldNotEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field cannot be empty"})
		return
	}

	c.JSON(201, services.InsertProduct(newProduct))
}

func PutProduct(c *gin.Context) {
	var newProduct models.Product
	id := c.Param("id")
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services.EditProduct(id, newProduct))
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(204, services.DeleteProduct(id))
}
