package mstore

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION_GOODS = "goods"
)

type MStore struct {
	db     *mongo.Database
	logger slog.Logger

	prodCache   map[cacheKey]prodCount
	prodCacheMu sync.Mutex
}

// Ключ для кэша счетчиков произведенных продуктов
type cacheKey struct {
	Gtin     string
	ProdDate time.Time
	Tname    string
}

// Значения кэша счетчиков
type prodCount struct {
	Produced  int64 // Количество произведенных
	Discarded int64 // Количество отбракованных
}

// Возвращает подключение к базе данных
func New(path string, dbname string, logger slog.Logger) (*MStore, error) {
	const op = "mstore.New"
	opts := options.Client().ApplyURI(path).SetTimeout(1000 * time.Millisecond)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверка подключения к базе
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	con := MStore{
		db:        client.Database(dbname),
		logger:    logger,
		prodCache: map[cacheKey]prodCount{},
	}

	return &con, nil
}
