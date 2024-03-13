package rest

import (
	"encoding/json"
	"fmt"
	"hub/internal/entity"
	"hub/internal/service/uadmin"

	"net/http"
)

// Добавляет продукт
// метод POST
func AddGood(u uadmin.UAdmin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		good_dto := good_dto{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good_dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGood := entity.Good{
			Gtin:            good_dto.Gtin,
			Desc:            good_dto.Desc,
			StoreCount:      good_dto.StoreCount,
			GetCodeForPrint: good_dto.GetCodeForPrint,
			AllowProduce:    good_dto.AllowProduce,
			AllowPrint:      good_dto.AllowPrint,
			Upload:          good_dto.Upload,
		}

		// Добавляем продукт в хранилище
		err = u.AddGood(r.Context(), mappedGood)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		resp_body := toResponse(true, "Успешно", nil)
		fmt.Fprint(w, resp_body)
	}
}

// Возвращает все продукты из базы
// метод POST
func GetAllGoods(u uadmin.UAdmin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Получаем продукты
		goods, err := u.GetAllGoods(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGoods := []good_dto{}
		for _, good := range goods {
			mappedGoods = append(mappedGoods, good_dto{
				Gtin:            string(good.Gtin),
				Desc:            good.Desc,
				StoreCount:      good.StoreCount,
				GetCodeForPrint: good.GetCodeForPrint,
				AllowProduce:    good.AllowProduce,
				AllowPrint:      good.AllowPrint,
				Upload:          good.Upload,
				CreatedAt:       good.CreatedAt,
			})
		}

		resp_body := toResponse(true, "Успешно", mappedGoods)
		fmt.Fprint(w, resp_body)
	}
}
