package mware

import (
	"hub/internal/api/http_web/authservice"
	"log"
	"net/http"
)

// Проверка, авторизован ли пользователь
func ChekAuth(auth *authservice.Auth, permission ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			// Отключена проверка авторизации
			// next.ServeHTTP(w, r)
			// return

			cookie, err := r.Cookie("authToken")
			if err != nil {
				// http.Error(w, authservice.ErrNotAuthorized.Error(), http.StatusUnauthorized)
				// Установка заголовков, чтобы браузер не кэшировал страницу
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Set("Pragma", "no-cache")
				w.Header().Set("Expires", "0")

				http.Redirect(w, r, "/login.html", http.StatusMovedPermanently)
				return
			}

			err = auth.Authorize(cookie.Value, permission)
			switch err {
			case authservice.ErrNotAuthorized:
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Set("Pragma", "no-cache")
				w.Header().Set("Expires", "0")

				http.Redirect(w, r, "/login.html", http.StatusMovedPermanently)

			case authservice.ErrInvalidCredentials:
				http.Error(w, err.Error(), http.StatusUnauthorized)

			case authservice.ErrNoPermission:
				http.Error(w, err.Error(), http.StatusUnauthorized)

			default:

			}

			if err != nil {
				log.Print(err)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
