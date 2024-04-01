package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_PRINTID = "printid"

type PrintIDCounter struct {
	Name  string `bson:"_id"`
	Value uint32
}

type CodeForPrint struct {
	Gtin    string
	Serial  string
	Crypto  string
	PrintID uint32
}

// Возвращает код для печати
func (m *MStore) GetCodeForPrint(ctx context.Context, gtin string, tname string, proddate string) (CodeForPrint, error) {
	// Логгирование
	const op = "mstore.GetCodeForPrint"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin)

	var err error
	var response CodeForPrint

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("response", response, "err", err, "duration", since)
		if err != nil || since > 10*time.Millisecond {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// - Проверить корректность входных данных
	if tname == "" {
		err = fmt.Errorf("не указано имя терминала")
		return CodeForPrint{}, err
	}

	err = entity.ValidateGtin(gtin)
	if err != nil {
		return CodeForPrint{}, err
	}

	// - Проверить корректность даты и преобразовать в time.Time
	_, err = time.Parse("2006-01-02", proddate) // YYYY-MM-DD
	if err != nil {
		return CodeForPrint{}, err
	}

	// Получаем код, пригодный для печати, ставим в бд флаг,
	// что он больше не доступен для печати, что бы заблокировать
	// возможность получения этого кода в другом потоке
	filter := bson.M{"printinfo.avaible": true}
	update := bson.M{"$set": bson.M{
		"printinfo.avaible":    false,
		"printinfo.tname":      tname,
		"printinfo.uploadtime": time.Now()},
	}

	var code entity.FullCode
	collect := m.db.Collection(gtin)
	err = collect.FindOneAndUpdate(ctx, filter, update).Decode(&code)
	if err != nil {
		return CodeForPrint{}, err
	}

	// Получаем для него printID из счетчика gtin + ":" + tname + ":" + дата фасовки
	// Инкрементируем счетчик кодов
	cname := gtin + ":" + tname + ":" + proddate // год месяц день
	filter = bson.M{"_id": cname}
	update = bson.M{"$inc": bson.M{"value": 1}}
	opt := options.FindOneAndUpdate().SetUpsert(true)

	var printID PrintIDCounter
	counters := m.db.Collection(COLLECTION_PRINTID)

	res := counters.FindOneAndUpdate(ctx, filter, update, opt)
	err = res.Decode(&printID)
	if err != nil {
		// Если этого счетчика раньше не было, то вернется ошибка
		// Запрашиваем обновление еще раз
		res = counters.FindOneAndUpdate(ctx, filter, update, opt)
		err = res.Decode(&printID)
		if err != nil {
			err = fmt.Errorf("ошибка инкремента счетчика %s", err)
			return CodeForPrint{}, err
		}
	}

	// Присваиваем коду PrintID
	filter = bson.M{"_id": code.Serial}
	update = bson.M{"$set": bson.M{"printinfo.printid": printID.Value}}
	updResult, err := m.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		err = fmt.Errorf("присввоение коду printId: %s", err)
		return CodeForPrint{}, err
	}

	// Проверяем, обновился ли id у кода
	if printID.Value > 0 && updResult.ModifiedCount != 1 {
		err = fmt.Errorf("ошибка установки printID для кода GTIN: %s serial: %s", gtin, code.Serial)
		return CodeForPrint{}, err
	}

	response = CodeForPrint{
		Gtin:    gtin,
		Serial:  code.Serial,
		Crypto:  code.Crypto,
		PrintID: printID.Value,
	}

	return response, nil
}
