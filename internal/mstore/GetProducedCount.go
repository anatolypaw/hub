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
func (m *MStore) GetProducedCount(ctx context.Context, tname string, gtin string, proddate string) (prodCount, error) {
	// Логгирование
	const op = "mstore.GetProducedCount"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("proddate", proddate)

	var err error
	var pcount prodCount

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("response", pcount, "err", err, "duration", since)
		if err != nil || since > 10*time.Millisecond {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// - Проверить корректность входных данных
	if tname == "" {
		err = fmt.Errorf("не указано имя терминала")
		return prodCount{-1, -1}, err
	}

	// - Проверить корректность gtin
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return prodCount{-1, -1}, err
	}

	// - Проверить корректность даты и преобразовать в time.Time
	tdate, err := time.Parse("2006-01-02", proddate) // YYYY-MM-DD
	if err != nil {
		return prodCount{-1, -1}, err
	}

	// Запрос в кэш бд за произведенными и отбракованными
	key := cacheKey{
		Gtin:     gtin,
		ProdDate: tdate,
		Tname:    tname,
	}

	m.prodCacheMu.Lock()
	pcount, ok := m.prodCache[key]
	m.prodCacheMu.Unlock()
	if ok {
		return pcount, nil
	}

	// В кэше нет данных, запрашиваем в бд количество произведенных
	// Произведенные те, у которых последнее событие - produce

	collection := m.db.Collection(gtin)

	// Поиск
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
		return prodCount{-1, -1}, err
	}

	// Получение результата агрегации
	var result []bson.M
	err = cursor.All(ctx, &result)
	if err != nil {
		return prodCount{-1, -1}, err
	}

	// Читаем результат
	if len(result) == 1 {
		// Получение значения "produced" из текущего элемента результата
		produced, ok := result[0]["produced"].(int32)
		if !ok {
			// Обработка ошибки, если приведение типа не удалось
			err = fmt.Errorf("ошибка приведения типа produced")
			return prodCount{-1, -1}, err
		}

		pcount.Produced = int64(produced)
	}

	// Подсчет отбракованных
	pipeline = mongo.Pipeline{
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
			"last.type":     "discard",
			"last.tname":    tname,
			"last.proddate": tdate,
		}}},
		{{Key: "$count", Value: "produced"}},
	}

	// Запускаем агрегацию
	cursor, err = collection.Aggregate(ctx, pipeline)
	if err != nil {
		return prodCount{-1, -1}, err
	}

	// Получение результата агрегации
	err = cursor.All(ctx, &result)
	if err != nil {
		return prodCount{-1, -1}, err
	}

	// Читаем результат
	if len(result) == 1 {
		// Получение значения "discard" из текущего элемента результата
		discard, ok := result[0]["produced"].(int32)
		if !ok {
			// Обработка ошибки, если приведение типа не удалось
			err = fmt.Errorf("ошибка приведения типа produced")
			return prodCount{-1, -1}, err
		}

		pcount.Discarded = int64(discard)
	}

	// Обновляем кэш
	m.prodCacheMu.Lock()
	m.prodCache[key] = pcount
	m.prodCacheMu.Unlock()

	return pcount, err
}
