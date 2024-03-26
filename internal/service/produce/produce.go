package produce

import (
	"context"
	"errors"
	"fmt"
	"hub/internal/entity"
	"log/slog"
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
