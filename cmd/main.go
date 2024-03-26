package main

import (
	grpcapi "hub/internal/api/grpc"
	pb "hub/internal/api/grpc/grpcapi"
	"hub/internal/mstore"
	"net"

	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
)

func main() {
	/* Настройка логгера */
	//logger := slog.New(slog.Default().Handler())
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	/* Подключение к базе данных */
	mstore, err := mstore.New("mongodb://localhost:27017/", "molocode", *logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	/* Инициализация http сервера */
	router := chi.NewRouter()

	// Admin
	// router.Post("/v1/admin/addGood", rest.AddGood(uadmin))
	// router.Get("/v1/admin/getAllGoods", rest.GetAllGoods(uadmin))

	// Exchange
	// router.Get("/v1/exchange/getGoodsReqCodes", rest.GetGoodsReqCodes(uexhange))
	// router.Post("/v1/exchange/addCodeForPrint", rest.AddCodeForPrint(uexhange))

	httpserver := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Info("HTTP server run on", "addres", httpserver.Addr)
		logger.Error(httpserver.ListenAndServe().Error())
	}()

	/* Инициализация gRPC сервера */
	lis, err := net.Listen("tcp", ":3100")
	if err != nil {
		logger.Error("failed to listen gRPC: %v", err)
	}

	var opts []grpc.ServerOption

	grpcserver := grpc.NewServer(opts...)
	grpcService := grpcapi.New(mstore)

	pb.RegisterHubServer(grpcserver, &grpcService)
	logger.Info("gRPC server run on", "addres", lis.Addr())
	if err := grpcserver.Serve(lis); err != nil {
		logger.Error("failed to serve gRPC: %v", err)
	}
}
