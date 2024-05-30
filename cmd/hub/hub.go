package main

import (
	"flag"
	grpcapi "hub/internal/api/grpc"
	pb "hub/internal/api/grpc/grpcapi"
	"hub/internal/config"
	"hub/internal/mstore"
	"hub/internal/web"
	"log"
	"net"

	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
)

const version = "2.0.2"

func main() {
	// Парсим флаги командной строки
	newConfigFlag := flag.Bool("new-config", false, "создать hub.json конфигурации по умолчанию.")
	flag.Parse()

	/* Настройка логгера */
	//logger := slog.New(slog.Default().Handler())
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	logger.Info("version", "version", version)

	/* Чтение настроек */
	// Создаем конфиг
	cfg := config.New("hub.json")
	// Если указан параметр, создаем файл конфигурации по умолчанию
	if *newConfigFlag {
		cfg.P = config.DefaultConfig
		err := cfg.Save()
		if err != nil {
			log.Print("Ошибка при создании файла конфигурации:", err)
			return
		}
		log.Print("Создан файл конфигурации по умолчанию ", "hub.cfg")
		return
	}

	err := cfg.Load()
	if err != nil {
		logger.Error("Загрузка конфигурации", err)
		os.Exit(1)
	}

	/* Запускаем web интерфейс */
	webui := web.New()
	go func() {
		err := webui.Run(":80")
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	/* Подключение к базе данных */
	mstore, err := mstore.New(cfg.P.MongoUri, cfg.P.DbName, *logger)
	if err != nil {
		logger.Error(err.Error())
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
