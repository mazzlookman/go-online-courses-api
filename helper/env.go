package helper

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
