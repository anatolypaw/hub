package produce

import (
	"context"
	"errors"
	"fmt"
	"hub/internal/entity"
	"log/slog"
	"time"
)

type iGoodRepo interface {
	Add(context.Context, entity.Good) error
	GetGood(context.Context, string) (entity.Good, error)
}

type iCodeRepo interface {
	AddCode(context.Context, entity.FullCode) error
	GetCode(ctx context.Context, gtin string, serial string) (entity.FullCode, error)
	GetCodeForPrint(ctx context.Context, gtin string, terminal string) (entity.Code, error)
}

type Produce struct {
	goodRepo iGoodRepo
	codeRepo iCodeRepo
	logger   slog.Logger
}

func New(goodRepo iGoodRepo, codeRepo iCodeRepo, logger slog.Logger) Produce {
	return Produce{
		goodRepo: goodRepo,
		codeRepo: codeRepo,
		logger:   logger,
	}
}

// Возвращает код для печати
func (u *Produce) GetCodeForPrint(ctx context.Context, gtin string, tname string) (entity.Code, error) {
	const op = "Produce.GetCodeForPrint"
	logger := u.logger.With("func", op).
		With("tname", tname).
		With("gtin", gtin)

	var err error
	var response entity.Code

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("response", response, "err", err, "duration", since)
		if err != nil {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// - Проверить корректность gtin
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return entity.Code{}, err
	}

	// - Проверить, разрешено ли для этого продукта выдача кодов для нанесения
	good, err := u.goodRepo.GetGood(ctx, gtin)
	if err != nil {
		return entity.Code{}, fmt.Errorf("ошибка запроса продукта: %s", err)
	}

	if !good.AllowPrint {
		return entity.Code{}, errors.New("для этого продукта запрещено выдача кодов для нанесения")
	}

	// - Получить код для печати
	// - TODO Проверить корректность кода в ответе БД
	response, err = u.codeRepo.GetCodeForPrint(ctx, gtin, tname)
	if err != nil {
		return entity.Code{}, err
	}

	return response, nil
}

// Отмечает ранее напечатанный код произведенным
func (usecase *Produce) ProducePrinted(ctx context.Context, gtin string, serial string, tname string, prodDate string) error {

	// - Проверить корректность gtin
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	// - Проверить корректность serial
	err = entity.ValidateSerial(serial)
	if err != nil {
		return err
	}

	// - Проверить корректность даты

	// - Проверить корректность имени терминала

	// - Проверить, разрешено ли производство для этого продукта
	good, err := usecase.goodRepo.GetGood(ctx, gtin)
	if err != nil {
		return fmt.Errorf("ошибка запроса продукта: %s", err)
	}

	if !good.AllowProduce {
		return errors.New("для этого продукта запрещено производство")
	}

	// - Проверки статуса кода
	code, err := usecase.codeRepo.GetCode(ctx, gtin, serial)

	// Проверить, был ли этот код отправлен на печать
	if code.PrintInfo.Avaible {
		return errors.New("этот код не был передан на печать")
	}

	// - Проверить, со
	panic(err)
}
