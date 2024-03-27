package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Отмечает код отбракованным
func (m *MStore) Discard(ctx context.Context, tname string, gtin string, serial string) error {
	// Логгирование
	const op = "mstore.Discard"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("serial", serial)

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

	// Проверка, что код уже произведен
	if !code.Produced {
		err = fmt.Errorf("код не произведен")
		return err
	}

	// Отмечаем его отбракованным
	filter = bson.M{"_id": serial}
	update := bson.M{"$set": bson.M{
		"produced":   false,
		"proddate":   0,
		"prodtime":   time.Now(),
		"prodtname":  tname,
		"discard":    false,
		"needupload": true,
	},
	}
	_, err = m.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Оновляем данные в кэше, увеличиваем количество произведенных
	key := cacheKey{
		Gtin:     gtin,
		ProdDate: tdate,
		Tname:    tname,
	}

	m.prodCacheMu.Lock()
	value, ok := m.prodCache[key]
	// Обновляем счетчики, только если этот ключ был в кэше
	// иначе счет пойдет с 0
	if ok {
		m.prodCache[key] = prodCount{
			Produced:  value.Produced + 1,
			Discarded: value.Discarded,
		}
	}
	m.prodCacheMu.Unlock()
}