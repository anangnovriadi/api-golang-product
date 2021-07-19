package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/anangnovriadi/api-golang-product/models"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.Connect()

	r := Router()
	r.Run(":" + os.Getenv("PORT"))
}
