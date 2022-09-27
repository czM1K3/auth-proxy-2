package auth

import (
	"time"

	"github.com/czM1K3/auth-proxy-2/src/env"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func generateJwt(expirationTime time.Time) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(env.GetSecret())
	return tokenString, err
}

func validateJwt(tknStr string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return env.GetSecret(), nil
	})
	if err != nil {
		return false
	}
	return tkn.Valid
}
