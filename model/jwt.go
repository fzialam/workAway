package model

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("233f42ec-f1c9-4558-befa-eff4292ea94b")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
