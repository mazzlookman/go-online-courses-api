package main

import (
	"fmt"
	"go-pzn-restful-api/app"
)

func main() {
	db := app.DBConnection()
	err := app.DBMigrate(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("migrate success")
}
