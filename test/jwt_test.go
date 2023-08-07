package test

import (
	"go-pzn-restful-api/auth"
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwtAuth := auth.NewJwtAuthImpl()
	jwtToken, err := jwtAuth.GenerateJwtToken(10)
	if err != nil {
		log.Println(err)
	}

	log.Println(jwtToken)
	log.Println(len(jwtToken))
}
