package rest

import (
	"hub/internal/domain/usecase"
	"net/http"
)

type server struct {
	httpServer *http.Server
	hub        *usecase.Hub
}

func New(hub *usecase.Hub) server {
	return server{hub: hub}
}

func (s *server) Run(addr string) error {
	mux := http.NewServeMux()

	// Добавляет продукт
	mux.HandleFunc("POST /api/goods/new", addGoodHandler(s.hub))

	// Возвращает продукт по запрошенному gtin
	mux.HandleFunc("GET /api/goods/{gtin}", getGoodHandler(s.hub))

	// Возвращает все продукты
	mux.HandleFunc("GET /api/goods/", getAllGoodsHandler(s.hub))

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return s.httpServer.ListenAndServe()
}
