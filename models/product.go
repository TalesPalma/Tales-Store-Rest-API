package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
	Quantity int    `json:"quantity"`
}
