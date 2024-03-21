package main

import (
	"fmt"
	"net/http"
)

const portNum = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the homepage")
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
