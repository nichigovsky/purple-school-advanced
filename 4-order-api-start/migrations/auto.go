package main

import (
	"4-order-api/internal/products"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&products.Product{})
}
