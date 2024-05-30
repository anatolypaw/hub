package mware

import (
	"hub/internal/web/authservice"
	"log"
	"net/http"
)

// Проверка, авторизован ли пользователь
func ChekAuth(auth *authservice.Auth, permission ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				// http.Error(w, authservice.ErrNotAuthorized.Error(), http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
				return
			}

			err = auth.Authorize(cookie.Value, permission)
			switch err {
			case authservice.ErrNotAuthorized:
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)

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
