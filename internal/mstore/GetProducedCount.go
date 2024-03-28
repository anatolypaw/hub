package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Выводит количество фасованных и отбракованных кодов дя запрошенной линии, продукта и  даты
func (m *MStore) GetProducedCount(ctx context.Context, tname string, gtin string, proddate string) (int64, error) {
	// Логгирование
	const op = "mstore.GetProducedCount"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("proddate", proddate)

	var err error
	var thisTerm int64

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("thisTerm", thisTerm, "err", err, "duration", since)
		if err != nil || since > 10*time.Millisecond {
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

	// Запрос в кэш бд за произведенными и отбракованными
	key := cacheKeyProdOnTerm{
		Gtin:     gtin,
		ProdDate: tdate,
		Tname:    tname,
	}

	m.CacheProdOnTermMu.Lock()
	thisTerm, ok := m.cacheProdOnTerm[key]
	m.CacheProdOnTermMu.Unlock()
	if ok {
		return thisTerm, nil
	}

	// В кэше нет данных, запрашиваем в бд количество произведенных
	// Произведенные те, у которых последнее событие - produce

	collection := m.db.Collection(gtin)

	// Подсчет произведенных кодов на этой линии
	// Определяем агрегационный конвейер
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"prodinfo.tname":    tname,
			"prodinfo.proddate": tdate,
		}}},
		{{Key: "$set", Value: bson.D{
			{Key: "last", Value: bson.D{
				{Key: "$last", Value: "$prodinfo"},
			}},
		}}},
		{{Key: "$match", Value: bson.M{
			"last.type":     "produce",
			"last.tname":    tname,
			"last.proddate": tdate,
		}}},
		{{Key: "$count", Value: "produced"}},
	}

	// Запускаем агрегацию
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return -1, err
	}

	// Получение результата агрегации
	var event []bson.M
	err = cursor.All(ctx, &event)
	if err != nil {
		return -1, err
	}

	// Читаем результат
	if len(event) == 1 {
		// Получение значения "produced" из текущего элемента результата
		produced, ok := event[0]["produced"].(int32)
		if !ok {
			// Обработка ошибки, если приведение типа не удалось
			err = fmt.Errorf("ошибка приведения типа produced")
			return -1, err
		}

		thisTerm = int64(produced)
	}

	// Обновляем кэш
	m.CacheProdOnTermMu.Lock()
	m.cacheProdOnTerm[key] = thisTerm
	m.CacheProdOnTermMu.Unlock()

	return thisTerm, err
}
