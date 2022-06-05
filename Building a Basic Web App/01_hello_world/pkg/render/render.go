package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate renders templates using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	pasredTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := pasredTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
