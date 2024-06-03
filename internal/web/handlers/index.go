package handlers

import (
	"html/template"
	"net/http"
)

var indexTemplate = template.Must(template.ParseFS(templates, "templates/index.html"))

// Форма аутентификации
func Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}
