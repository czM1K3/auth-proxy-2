package main

import (
	"log"
	"net/http"

	"github.com/czM1K3/auth-proxy-2/src/env"
	"github.com/czM1K3/auth-proxy-2/src/proxy"
)

func main() {
	env.GetPassword()
	env.GetLoginTime()

	port := "4000"
	http.HandleFunc("/", proxy.HandleRequest)

	log.Println("Starting proxy on http://localhost:" + port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
