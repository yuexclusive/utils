package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenToken(id string, key string) (string, error) {
	if id == "" {
		return "", errors.New("invalid id")
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "future",
		Id:        id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetClaims(token string, key string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	jt, err := jwt.ParseWithClaims(
		token, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !jt.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
