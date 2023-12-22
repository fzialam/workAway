package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var JWT_KEY = []byte("e938d81d-cc6c-43aa-a5e8-1ed9b0a88566")

type JWTClaims struct {
	Id   int `json:"id"`
	Role int `json:"role"`
	jwt.RegisteredClaims
}

func VerifyRoleUser(role int) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("jwt-workaway")
			if err != nil {
				err = errors.New("Token tidak valid")
				helper.UnauthorizedTemplate(w, r, err.Error())
				return
			}

			tokenCookie := c.Value
			claims := &JWTClaims{}

			token, err := jwt.ParseWithClaims(tokenCookie, claims, func(t *jwt.Token) (interface{}, error) {
				return JWT_KEY, nil
			})

			if err != nil || !token.Valid || claims.Role != role {
				if !token.Valid {
					err = errors.New("Token tidak valid")
				} else if claims.Role != role {
					err = errors.New("URL terlindungi")
				}
				helper.UnauthorizedTemplate(w, r, err.Error())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func LoginSetJWT(user entity.User) string {
	expTime := time.Now().Add(30 * time.Minute)
	claims := JWTClaims{
		Id:   user.Id,
		Role: user.Rank,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-workaway",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(JWT_KEY)
	helper.PanicIfError(err)

	return token
}
