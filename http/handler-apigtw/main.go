package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	homeServiceURL, _ := url.Parse("http://localhost:8081")
	aboutServiceURL, _ := url.Parse("http://localhost:8082")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/about" {
			proxy := httputil.NewSingleHostReverseProxy(aboutServiceURL)
			proxy.ServeHTTP(w, r)
		} else {
			proxy := httputil.NewSingleHostReverseProxy(homeServiceURL)
			proxy.ServeHTTP(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
