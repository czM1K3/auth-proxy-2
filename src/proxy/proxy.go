package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/czM1K3/auth-proxy-2/src/auth"
	"github.com/czM1K3/auth-proxy-2/src/env"
)

const loginUrl = "/login"

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	// Check if token is valid
	if auth.Check(req) {
		// Proxy address
		serveReverseProxy(res, req)
	} else {
		// Check if is on login page
		if strings.HasSuffix(req.URL.Path, loginUrl) {
			// If so, check for request method
			if req.Method == "POST" {
				// If POST, try to login
				if auth.Login(res, req) {
					// If password correct, generate token, save it as cookie and redirect to index page
					http.Redirect(res, req, "/", http.StatusSeeOther)
				} else {
					// If password was not right, show incorrect message
					http.Redirect(res, req, loginUrl+"?success=false", http.StatusSeeOther)
				}
			} else {
				// If not, show login page
				data, error := os.ReadFile("public/login.html")
				if error != nil {
					data = []byte("Login page error")
				}
				res.Write(data)
			}
		} else {
			// Else redirect to login page
			http.Redirect(res, req, loginUrl, http.StatusTemporaryRedirect)
		}
	}
}

func serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(env.GetServiceAddress())
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = errorHandler
	proxy.ServeHTTP(res, req)
}

func errorHandler(res http.ResponseWriter, _ *http.Request, _ error) {
	res.Write([]byte("Proxy error"))
}
