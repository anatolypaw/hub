package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Отмечает напечатанный код произведенным
func (m *MStore) ProducePrinted(ctx context.Context, tname string, gtin string, serial string, proddate string) error {
	// Логгирование
	const op = "mstore.ProducePrinted"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("serial", serial)

	var err error

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("err", err, "duration", since)
		if err != nil {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// - Проверить корректность входных данных
	if tname == "" {
		err = fmt.Errorf("не указано имя терминала")
		return err
	}

	// - Проверить корректность gtin
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	// - Проверить корректность serial
	err = entity.ValidateSerial(serial)
	if err != nil {
		return err
	}

	// - Проверить корректность даты и преобразовать в time.Time
	tdate, err := time.Parse("2006-01-02", proddate) // YYYY-MM-DD
	if err != nil {
		return err
	}

	// Получить код и бд
	filter := bson.M{"_id": serial}
	collect := m.db.Collection(gtin)
	reqResult := collect.FindOne(ctx, filter)

	var code entity.FullCode
	err = reqResult.Decode(&code)
	if err != nil {
		err = fmt.Errorf("код не найден: %s", err)
		return err
	}

	if code.Serial == "" {
		err = fmt.Errorf("код не найден")
		return err
	}

	// Проверка, что код не был уже произведен
	if code.Produced {
		err = fmt.Errorf("код уже произведен %s", code.ProdTime)
		return err
	}

	// Добавляем данные о производстве
	filter = bson.M{"_id": serial}
	update := bson.M{"$set": bson.M{
		"produced":  true,
		"proddate":  tdate,
		"prodtime":  time.Now(),
		"prodtname": tname,
		"discard":   false,
		"uploaded":  false,
	},
	}
	_, err = m.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
