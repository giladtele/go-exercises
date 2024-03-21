package main

import (
	"fmt"
	"net/http"
)

const portNum = ":8082"

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the about page")
}

func main() {
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
