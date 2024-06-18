package handlers

import (
	"html/template"
	"net/http"
)

var debugTemplate = template.Must(template.ParseFS(templates, "templates/debug.html"))

// Форма аутентификации
func Debug(w http.ResponseWriter, r *http.Request) {
	debugTemplate.Execute(w, nil)
}
