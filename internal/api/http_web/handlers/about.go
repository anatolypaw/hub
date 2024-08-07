package handlers

import (
	"encoding/json"
	"net/http"
)

// Возвращает текущую версию ПО
func AboutInfo(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Resp struct {
			Version string `json:"version"`
		}

		resp := Resp{
			Version: version,
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
