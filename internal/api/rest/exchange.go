package rest

import (
	"encoding/json"
	"fmt"
	"hub/internal/entity"
	"hub/internal/service/uexchange"

	"net/http"
)

// Добавляет продукт
// метод POST
func GetGoodsReqCodes(u uexchange.UExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Получаем список продуктов и требуемое количество кодов
		codereq, err := u.GetGoodsReqCodes(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		type codeReq_json struct {
			Gtin     string `json:"gtin"`
			Desc     string `json:"desc"`
			Required uint   `json:"required"`
		}

		// MAPPING
		mappedCodeReq := []codeReq_json{}
		for _, ths := range codereq {
			mappedCodeReq = append(mappedCodeReq, codeReq_json{
				Gtin:     ths.Gtin,
				Desc:     ths.Desc,
				Required: ths.Required,
			})
		}

		resp_body := toResponse(true, "Успешно", mappedCodeReq)
		fmt.Fprint(w, resp_body)
	}
}

// Доабвляет код для печати
func AddCodeForPrint(u uexchange.UExchange) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// - Получаем код из body
		type AddCode_json struct {
			Gtin       string `json:"gtin"`
			Serial     string `json:"serial"`
			Сrypto     string `json:"crypto"`
			SourceName string `json:"source_name"`
		}

		code_json := AddCode_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&code_json)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedCode := entity.Code{
			Gtin:   code_json.Gtin,
			Serial: code_json.Serial,
			Crypto: code_json.Сrypto,
		}
		// Добавляем продукт
		err = u.AddCodeForPrint(r.Context(), mappedCode, code_json.SourceName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		resp_body := toResponse(true, "Успешно", nil)
		fmt.Fprint(w, resp_body)
	}
}
