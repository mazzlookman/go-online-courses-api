package app

import (
	"context"
	"go-pzn-restful-api/helper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
)

func NgrokInit() ngrok.Tunnel {
	tun, err := ngrok.Listen(context.Background(),
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	helper.PanicIfError(err)
	log.Println("Ngrok URL: ", tun.URL())

	return tun
}
