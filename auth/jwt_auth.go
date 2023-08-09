package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("inirahasiabanget")

type JwtAuth interface {
	GenerateJwtToken(role string, id int) (string, error)
	ValidateJwtToken(token string) (*jwt.Token, error)
}

type JwtAuthImpl struct {
}

func NewJwtAuth() JwtAuth {
	return &JwtAuthImpl{}
}

func (j *JwtAuthImpl) GenerateJwtToken(role string, id int) (string, error) {
	mapClaims := jwt.MapClaims{}

	if role == "author" {
		mapClaims["author_id"] = id
	}

	if role == "user" {
		mapClaims["user_id"] = id
	}

	tokenWithHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	signedToken, err := tokenWithHeader.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JwtAuthImpl) ValidateJwtToken(token string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		return parseToken, err
	}

	return parseToken, nil
}
