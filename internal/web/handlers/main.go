package handlers

import (
	"html/template"
	"net/http"
)

var mainTemplate = template.Must(template.ParseFiles("templates/main.html"))

// Форма аутентификации
func MainForm(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w, nil)
}
