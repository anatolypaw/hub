package mstore

import (
	"context"
	"errors"
	"fmt"
	"hub/internal/entity"
	"time"
)

// Добавлет код для печати в бд
func (m *MStore) AddCodeForPrint(ctx context.Context, sname string, gtin string, serial string, crypto string) error {
	// Логгирование
	const op = "mstore.AddCodeForPrint"
	logger := m.logger.With("func", op).
		With("sname", sname).
		With("gtin", gtin).
		With("serial", serial).
		With("crypto", crypto)

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

	// - Проверить корректность кода
	if sname == "" {
		err = fmt.Errorf("не указано имя источника")
		return err
	}
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	err = entity.ValidateSerial(serial)
	if err != nil {
		return err
	}

	err = entity.ValidateCrypto(crypto)
	if err != nil {
		return err
	}

	// - Проверить, разрешено ли для этого продукта добавление кодов
	good, err := m.GetGood(ctx, gtin)
	if err != nil {
		return err
	}

	if !good.GetCodeForPrint {
		return errors.New("для этотого продукта запрещено получение кодов")
	}

	// - Добавить код для печати
	code := entity.FullCode{
		Serial: serial,
		Crypto: crypto,
		Type:   "print",
		PrintInfo: entity.PrintInfo{
			Sname:   sname,
			Loaded:  time.Now(),
			Avaible: true,
		},
	}

	collect := m.db.Collection(gtin)
	_, err = collect.InsertOne(ctx, code)
	if err != nil {
		return err
	}
	return err
}
