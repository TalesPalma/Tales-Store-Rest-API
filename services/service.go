package services

import (
	"net/http"

	"github.com/TalesPalma/gin-golang-rest/database"
	"github.com/TalesPalma/gin-golang-rest/models"
)

func FindProducts() []models.Product {
	var products []models.Product
	database.DB.Find(&products)
	return products
}

func FindProductsById(id string) models.Product {
	var product models.Product
	database.DB.Find(&product, id)
	return product
}

func InsertProduct(product models.Product) models.Product {
	database.DB.Create(&product)
	return product
}

func EditProduct(id string, product models.Product) models.Product {
	database.DB.Model(&product).Where("id = ?", id).Updates(&product)
	return product
}

func DeleteProduct(id string) (string, int) {
	var product models.Product
	err := database.DB.Find(&product, id)
	if err.RowsAffected == 0 {
		return "Not found product", http.StatusNotFound
	}
	if err := database.DB.Delete(&product, id); err.Error != nil {
		return "Error to delete product", http.StatusInternalServerError
	}
	return "Product deleted successfully", http.StatusOK
}
