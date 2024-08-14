package controllers

import (
	"net/http"

	"github.com/TalesPalma/gin-golang-rest/models"
	"github.com/TalesPalma/gin-golang-rest/services"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, services.GetProducts())
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(
		200,
		services.GetProductsById(id),
	)
}

func PostProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, services.PostProduct(newProduct))
}
