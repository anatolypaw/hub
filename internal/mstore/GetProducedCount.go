package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Выводит количество фасованных кодов дя запрошенной линии, продукта, даты
func (m *MStore) GetProducedCount(ctx context.Context, tname string, gtin string, proddate string) (int64, error) {
	// Логгирование
	const op = "mstore.GetProducedCount"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("proddate", proddate)

	var err error
	var response int

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("response", response, "err", err, "duration", since)
		if err != nil {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// - Проверить корректность входных данных
	if tname == "" {
		err = fmt.Errorf("не указано имя терминала")
		return -1, err
	}

	// - Проверить корректность gtin
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return -1, err
	}

	// - Проверить корректность даты и преобразовать в time.Time
	tdate, err := time.Parse("2006-01-02", proddate) // YYYY-MM-DD
	if err != nil {
		return -1, err
	}

	// Запрос в бд
	filter := bson.M{
		"produced":  true,
		"proddate":  tdate,
		"prodtname": tname,
	}

	collect := m.db.Collection(gtin)
	count, err := collect.CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}

	return count, err
}
