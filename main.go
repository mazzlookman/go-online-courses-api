package main

import (
	"go-pzn-restful-api/app"
)

func main() {
	app.EnvInit()
	listener := app.NgrokInit()

	router := app.NewRouter()
	router.RunListener(listener)
	router.SetTrustedProxies([]string{listener.URL()})

	router.Run()
}
