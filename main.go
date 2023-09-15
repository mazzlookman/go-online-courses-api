package main

import (
	"go-pzn-restful-api/app"
)

func main() {
	router := app.NewRouter()
	router.Run(":2802")
}
