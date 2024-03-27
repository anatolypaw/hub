package main

import (
	grpcapi "hub/internal/api/grpc"
	pb "hub/internal/api/grpc/grpcapi"
	"hub/internal/mstore"
	"net"

	"log/slog"
	"os"

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
