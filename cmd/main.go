package main

import (
	"hub/internal/ctxlogger"
	"hub/internal/mstore"
	"hub/internal/rest"
	"hub/internal/usecase/uadmin"
	"hub/internal/usecase/uexchange"
	"hub/internal/usecase/uproduce"

	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/lmittmann/tint"
)

func main() {
	/* Настройка логгера */
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	/* Подключение к базе данных */
	mstore, err := mstore.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	/* Инициализация usecase */
	uadmin := uadmin.New(mstore)
	uexhange := uexchange.New(mstore, mstore)
	uproduce := uproduce.New(mstore, mstore)

	/* Инициализация http сервера */
	router := chi.NewRouter()

	// Логгер slog встраивается в context
	// на каждый request создается уникальный req_id и встраивается в context
	// он выводится в лог для всего дерева вызовов
	router.Use(ctxlogger.Logger(logger))

	// Admin
	router.Post("/v1/admin/addGood", rest.AddGood(uadmin))
	router.Get("/v1/admin/getAllGoods", rest.GetAllGoods(uadmin))

	// Exchange
	router.Get("/v1/exchange/getGoodsReqCodes", rest.GetGoodsReqCodes(uexhange))
	router.Post("/v1/exchange/addCodeForPrint", rest.AddCodeForPrint(uexhange))

	// Produce
	router.Get("/v1/produce/getCodeForPrint", rest.GetCodeForPrint(uproduce))
	router.Get("/v1/produce/producePrinted", rest.ProducePrinted(uproduce))

	s := &http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Info("Server run on", "addres", s.Addr)
	logger.Error(s.ListenAndServe().Error())
}
