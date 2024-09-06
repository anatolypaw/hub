package rest

import (
	"encoding/json"
	"fmt"
	"hub/internal/domain/models"
	"hub/internal/domain/usecase"
	"io"
	"net/http"
)

// POST /api/goods/new
func addGoodHandler(hub *usecase.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var good models.Good
		if err := json.Unmarshal(body, &good); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if err = hub.NewGood(good); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Fprint(w, "ok")
	}
}

// GET /api/goods/{gtin}/
func getGoodHandler(hub *usecase.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gtin := r.PathValue("gtin")

		good, err := hub.GetGood(gtin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(good)
	}

}

// GET /api/goods/
func getAllGoodsHandler(hub *usecase.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		goods, err := hub.GetAllGoods()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(goods)
	}
}
