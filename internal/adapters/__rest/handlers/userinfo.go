package handlers

import (
	"encoding/json"
	"hub/internal/api/http_web/authservice"
	"net/http"
)

// Возвращает текущее имя пользователя, определяет по кукам
func UserInfo(auth *authservice.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Resp struct {
			Username string `json:"username"`
		}

		cookie, err := r.Cookie("authToken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: "))
			w.Write([]byte(err.Error()))
			return
		}

		resp := Resp{
			Username: auth.GetUsernameByAuthToken(cookie.Value),
		}
		m, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: "))
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(m)
	}
}
