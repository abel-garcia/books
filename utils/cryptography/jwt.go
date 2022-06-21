package cryptography

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	exp_time   = 10
	signingKey = "3XDr8VwE<4P4&G"
)

type CustomClaims struct {
	Roles string `json:"roles,omitempty"`
	*jwt.StandardClaims
}

func CreateJWTWithClaims(userID string) (string, error) {
	claims := CustomClaims{
		Roles: "admin",
		StandardClaims: &jwt.StandardClaims{
			Issuer:    "books",
			Audience:  "user",
			Subject:   userID,
			ExpiresAt: time.Now().Add(exp_time * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func DecodeJWTSigned(token string) (*jwt.Token, *CustomClaims, error) {
	claims := &CustomClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	return jwtToken, claims, nil
}
