package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Устанавливает запись производства как выгруженную
func (m *MStore) SetCodeUploaded(ctx context.Context, gtin string, serial string, entryid string) error {
	// Логгирование
	const op = "mstore.GetCodeForPrint"
	logger := m.logger.With("func", op).
		With("gtin", gtin).
		With("serail", serial).
		With("id", entryid)

	var err error

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("err", err, "duration", since)
		if err != nil || since > 10*time.Millisecond {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	err = entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	// Обновляем запись
	// Условие для обновления записи
	filter := bson.M{"_id": serial, "prodinfo.id": entryid}
	update := bson.M{"$set": bson.M{
		"prodinfo.$.uploaded":   true,
		"prodinfo.$.uploadtime": time.Now(),
	}}

	collect := m.db.Collection(gtin)
	result, err := collect.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Проверить, было ли обновление успешным
	if result.ModifiedCount == 1 {
		return nil
	} else {
		return fmt.Errorf("измненено записей %d", result.MatchedCount)
	}
}
