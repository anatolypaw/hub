package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Отмечает код отбракованным
func (m *MStore) DiscardBySerial(ctx context.Context, tname string, gtin string, serial string) error {
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
		err = fmt.Errorf("код не фасован: %s", err)
		return err
	}

	if code.Serial == "" {
		err = fmt.Errorf("код не найден")
		return err
	}

	// Проверяем, что есть записи в логе
	if len(code.ProdInfo) == 0 {
		err = fmt.Errorf("лог пустой, код не был произведен")
		return err
	}

	// Получаем последнюю запись
	last := code.ProdInfo[len(code.ProdInfo)-1]

	// Проверяем, что он не был ранее отбракован
	if last.Type == "discard" {
		err = fmt.Errorf("код уже отбракован")
		return err
	}

	// Проверяем, что он был произведен
	if last.Type != "produce" {
		err = fmt.Errorf("код не был произведен")
		return err
	}

	// Добавляем данные о его отбраковке в массив (лог)
	prodInfo := entity.ProdInfo{
		ID:   uuid.New().String(),
		Time: time.Now(),
		Type: "discard",
		// Записываем в отбраковку дату, которой код ранее был произведен, так как в шлюз при отбраковке нужно передавать дату, которой код был произведен
		ProdDate: last.ProdDate,
		Tname:    tname,
	}

	filter = bson.M{"_id": serial}
	update := bson.M{"$push": bson.M{
		"prodinfo": prodInfo,
	},
	}
	_, err = m.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Обновляем данные в кэше, увеличиваем количество произведенных
	// Уменьшаем счетчик на той линии, на которой код был произведен, независимо от того, где он был отбракован
	key := cacheKeyProdOnTerm{
		Gtin:     gtin,
		ProdDate: last.ProdDate,
		Tname:    last.Tname,
	}

	m.CacheProdOnTermMu.Lock()
	value, ok := m.cacheProdOnTerm[key]
	// Обновляем счетчики, только если этот ключ был в кэше
	// иначе счет пойдет с 0
	if ok {
		m.cacheProdOnTerm[key] = value - 1
	}
	m.CacheProdOnTermMu.Unlock()

	return nil
}
