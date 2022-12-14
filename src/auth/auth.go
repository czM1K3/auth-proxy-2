package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/czM1K3/auth-proxy-2/src/env"
)

const (
	cookieName = "ProxyAuthorization"
)

func Check(req *http.Request) bool {
	cookie, error := req.Cookie(cookieName)
	if error != nil {
		return false
	}
	return validateJwt(cookie.Value)
}

func Login(res http.ResponseWriter, req *http.Request) bool {
	req.ParseForm()
	value := req.FormValue("password")
	if env.GetPassword() == value {
		expirationTime := time.Now().Add(time.Duration(env.GetLoginTime()) * time.Minute)
		token, err := generateJwt(expirationTime)
		if err != nil {
			log.Println("Something went wrong when generating token")
			return false
		}

		http.SetCookie(res, &http.Cookie{
			Name:     cookieName,
			Value:    token,
			Expires:  expirationTime,
			Path:     "/",
			HttpOnly: true,
		})
		return true
	}
	return false
}
