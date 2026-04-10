package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"5-order-api-auth/internal/users"
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

	db.AutoMigrate(&users.User{})
}
