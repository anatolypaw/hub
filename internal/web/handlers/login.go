package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var loginTemplate = template.Must(template.ParseFS(templates, "templates/login.html"))

// Форма аутентификации
func LoginForm(w http.ResponseWriter, r *http.Request) {
	loginTemplate.Execute(w, nil)
}

// Сервис аутентификации
type IAuth interface {
	Login(username, password string) (token string, error error)
}

// Аутентификация, возвращает sessionid в браузер в случае успеха
func Login(auth IAuth) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		sessionID, err := auth.Login(username, password)
		if err != nil {
			w.Write([]byte("Неверный логин или пароль"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println("Пользователь", username, "авторизован", "токен", sessionID)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   sessionID,
			Expires: time.Now().Add(24 * time.Hour),
		})

		w.Header().Set("HX-Redirect", "/")
		//w.Write([]byte("<p>Login successful!</p>"))
		w.WriteHeader(http.StatusOK)
	}

	return http.HandlerFunc(fn)
}
