package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNum = ":8082"

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatalln("error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
