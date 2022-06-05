package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate just renders templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	pasredTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := pasredTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
