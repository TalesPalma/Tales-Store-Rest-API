package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TalesPalma/gin-golang-rest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDatabase() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", user, password, dbName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Teste", user, password, dbName)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Product{})
}
