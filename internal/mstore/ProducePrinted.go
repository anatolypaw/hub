package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"github.com/google/uuid"
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

	// Проверка, что код был напечтан
	if code.PrintInfo.Avaible {
		err = fmt.Errorf("код не напечатан")
		return err
	}

	// Проверка, что код уже произведен по данным в последнем элементе лога
	// Возможна ситуация, что терминал дважды попытается выгрузить код
	// Такие коды можно найти в базе с двумя подряд отметками о производстве.
	// Если вернуть ошибку на терминал, то терминал встанет в бесконечную выгрузку этого кода
	last := len(code.ProdInfo) - 1
	if last >= 0 {
		if code.ProdInfo[last].Type == "produce" {
			err = fmt.Errorf("код уже произведен")
			//	return err
		}
	}

	// Добавляем данные о производстве в массив (лог)
	prodInfo := entity.ProdInfo{
		ID:       uuid.New().String(),
		Time:     time.Now(),
		Type:     "produce",
		ProdDate: tdate,
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

	// Оновляем данные в кэше, увеличиваем количество произведенных
	key := cacheKeyProdOnTerm{
		Gtin:     gtin,
		ProdDate: tdate,
		Tname:    tname,
	}

	m.CacheProdOnTermMu.Lock()
	prodCount, ok := m.cacheProdOnTerm[key]
	// Обновляем счетчик, только если этот ключ был в кэше
	// иначе счет пойдет с 0
	if ok {
		m.cacheProdOnTerm[key] = prodCount + 1
	}
	m.CacheProdOnTermMu.Unlock()

	return nil
}
