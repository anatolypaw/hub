package handlers

import (
	"context"
	"html/template"
	"hub/internal/mstore"
	"net/http"
)

var goodsTemplate = template.Must(template.ParseFS(templates, "templates/layout.html", "templates/goods.html"))

// Список продуктов
func GoodsGet(mstore *mstore.MStore) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		goods, err := mstore.GetAllGoods(context.TODO())
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		_ = goods
		goodsTemplate.Execute(w, nil)
	}
	return http.HandlerFunc(fn)
}
