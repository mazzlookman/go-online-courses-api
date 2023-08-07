package helper

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	PanicIfError(err)
	return os.Getenv(key)
}
