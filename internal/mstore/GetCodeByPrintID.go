package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает код по его Print ID
func (m *MStore) GetCodeByPrintID(ctx context.Context, tname, gtin, proddate string, printid uint32) (entity.Code, error) {
	// Логгирование
	const op = "mstore.GetCodeByPrintID"
	logger := m.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin).
		With("proddate", proddate).
		With("printID", printid)

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
		return entity.Code{}, err
	}

	// - Проверить корректность gtin
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return entity.Code{}, err
	}

	// - Проверить корректность даты и преобразовать в time.Time
	tdate, err := time.Parse("2006-01-02", proddate) // YYYY-MM-DD
	if err != nil {
		return entity.Code{}, err
	}

	//TODO
	// Получить код и бд
	filter := bson.M{"printinfo.printid": printid, "printinfo.tname": tname, "prodinfo.proddate": tdate}
	collect := m.db.Collection(gtin)
	reqResult := collect.FindOne(ctx, filter)

	var code entity.FullCode
	err = reqResult.Decode(&code)
	if err != nil {
		err = fmt.Errorf("код не найден: %s", err)
		return entity.Code{}, err
	}

	if code.Serial == "" {
		err = fmt.Errorf("код не найден")
		return entity.Code{}, err
	}

	// Проверяем, что есть записи в логе
	if len(code.ProdInfo) == 0 {
		err = fmt.Errorf("лог пустой, код не был произведен")
		return entity.Code{}, err
	}
	c := entity.Code{
		Gtin:   gtin,
		Serial: code.Serial,
		Crypto: code.Crypto,
	}

	return c, nil
}
