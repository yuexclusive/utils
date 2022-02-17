package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func GenToken(claim StandardClaims, key string) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Audience:  claim.Audience,
		ExpiresAt: claim.ExpiresAt,
		Id:        claim.Id,
		IssuedAt:  claim.IssuedAt,
		Issuer:    claim.Issuer,
		NotBefore: claim.NotBefore,
		Subject:   claim.Subject,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(key))

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
