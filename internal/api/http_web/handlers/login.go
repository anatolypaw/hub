package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// json ответ на запрос авторизации
type resp struct {
	Message   string `json:"message"`
	AuthToken string `json:"authToken"`
}

// Сервис аутентификации
type IAuth interface {
	Login(username, password string) (token string, error error)
}

// Аутентификация, возвращает sessionid в браузер в случае успеха
func Login(auth IAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Считываем логин и пароль из запроса
		var c creds
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: "))
			w.Write([]byte(err.Error()))
			return
		}

		// Проверка логина и пароля
		authToken, err := auth.Login(c.Username, c.Password)
		if err != nil {
			m, err := json.Marshal(resp{Message: err.Error()})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(m)
			return
		}

		// Отправляем токен авторизации
		m, err := json.Marshal(resp{AuthToken: authToken})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: "))
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(m)

		log.Println("Пользователь", c.Username, "авторизован", "токен", authToken)
	}
}
