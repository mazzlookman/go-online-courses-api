package test

import (
	"log"
	"path/filepath"
	"testing"
)

func TestRelPath(t *testing.T) {
	base := "service/user_service_impl.go"
	target := "assets/images/avatars/h.jpg"
	rel, err := filepath.Rel(base, target)
	if err != nil {
		log.Println(err)
	}

	log.Println(rel)
}
