package mstore

import (
	"context"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает код запрошенного gtin, который нужно выгрузить
type CodeForUpload struct {
	EntryID  string // ID записи, из которой получена информация, что код надо выгрузить
	Discard  bool   // Флаг, что код надо отбраковать
	Gtin     string
	Serial   string
	Crypto   string
	Proddate string
}

func (m *MStore) GetCodeForUpload(ctx context.Context, gtin string) (CodeForUpload, error) {
	// Логгирование
	const op = "mstore.GetCodeForUpload"
	logger := m.logger.With("func", op).
		With("gtin", gtin)

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
		return CodeForUpload{}, err
	}

	// Получаем код, который еще не выгружен
	filter := bson.M{"prodinfo.uploaded": false}

	var code entity.FullCode
	collect := m.db.Collection(gtin)
	err = collect.FindOne(ctx, filter).Decode(&code)
	if err != nil {
		return CodeForUpload{}, err
	}

	// Выдаем первую запись, которая не выгружена
	for _, entry := range code.ProdInfo {
		if !entry.Uploaded {
			return CodeForUpload{
					EntryID:  entry.ID,
					Discard:  entry.Type != "produce",
					Gtin:     gtin,
					Serial:   code.Serial,
					Crypto:   code.Crypto,
					Proddate: entry.ProdDate.Format("2006-01-02"),
				},
				nil
		}
	}

	return CodeForUpload{}, nil
}
