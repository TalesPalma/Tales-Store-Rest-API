package services

import (
	"github.com/TalesPalma/gin-golang-rest/database"
	"github.com/TalesPalma/gin-golang-rest/models"
)

func GetProducts() []models.Product {
	return []models.Product{}
}

func GetProductsById(id string) string {
	return "Certo" + id
}

func PostProduct(product models.Product) models.Product {
	database.DB.Create(&product)
	return product
}
