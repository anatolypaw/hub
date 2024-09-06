package handlers

import (
	"context"
	"encoding/json"
	"hub/internal/entity"
	"net/http"
)

type IGoodStore interface {
	GetAllGoods(context.Context) ([]entity.Good, error)
}

// Возвращает все продукты из хранилища
func GetAllGoods(goodStore IGoodStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Resp struct {
			Error  string
			Result []entity.Good
		}

		var m []byte
		goods, err := goodStore.GetAllGoods(r.Context())
		if err != nil {
			m, _ = json.Marshal(Resp{
				Error: err.Error(),
			})
		} else {
			m, _ = json.Marshal(Resp{
				Result: goods,
			})
		}

		w.Write(m)
	}
}
