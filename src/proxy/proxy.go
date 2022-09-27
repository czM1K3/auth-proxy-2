package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/czM1K3/auth-proxy-2/src/auth"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	isLogin := auth.Check(req)
	if !isLogin {
		if strings.Contains(req.URL.Path, "/auth/login") {
			if req.Method == "POST" {
				if auth.Login(res, req) {
					http.Redirect(res, req, "/", http.StatusSeeOther)
				} else {
					http.Redirect(res, req, "/auth/login?success=false", http.StatusSeeOther)
				}
			} else {
				data, error := os.ReadFile("public/login.html")
				if error != nil {
					data = []byte("Login page error")
				}
				res.Write(data)
			}
		} else {
			http.Redirect(res, req, "/auth/login", http.StatusTemporaryRedirect)
		}
	} else {
		serveReverseProxy("http://localhost:5555", res, req)
	}
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = errorHandler
	proxy.ServeHTTP(res, req)
}

func errorHandler(res http.ResponseWriter, _ *http.Request, _ error) {
	res.Write([]byte("Proxy error"))
}
