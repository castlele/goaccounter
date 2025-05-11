package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Key string    `json:"key"`
	Exp time.Time `json:"exp"`
}

var secretKey = []byte("JWT_SECRET")

func GenerateJWT(userName string, exp time.Time) (*JWT, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userName,
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	return &JWT{
		Key: key,
		Exp: exp,
	}, nil
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return token, err
}
