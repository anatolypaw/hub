package uexchange

import (
	"context"
	"hub/internal/entity"
)

type iGoodRepo interface {
	GetGood(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type iCodeRepo interface {
	AddCode(context.Context, entity.FullCode) error
	GetCountPrintAvaible(context.Context, string) (uint, error)
}

type UExchange struct {
	goodRepo iGoodRepo
	codeRepo iCodeRepo
}

func New(goodRepo iGoodRepo, codeRepo iCodeRepo) UExchange {
	return UExchange{
		goodRepo: goodRepo,
		codeRepo: codeRepo,
	}
}

// Возвращаемая кейсом структура
type CodeReq struct {
	Gtin     string
	Desc     string
	Required uint
}

// Возвращает список продуктов, требующих наполнения кодами для печати
// и количество требуемых кодов
func (u *UExchange) GetGoodsReqCodes(ctx context.Context,
) ([]CodeReq, error) {
	// - Получить продукты
	allGoods, err := u.goodRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// - Выбрать те, для которых включено наполнения кодами
	var goodsAvaibleForPrint []entity.Good
	for _, good := range allGoods {
		if good.GetCodeForPrint {
			goodsAvaibleForPrint = append(goodsAvaibleForPrint, good)
		}
	}

	// - Для каждого продукта получить доступное количество кодов
	var codesReq []CodeReq
	for _, good := range goodsAvaibleForPrint {
		avaibleCount, err := u.codeRepo.GetCountPrintAvaible(ctx, good.Gtin)
		if err != nil {
			return nil, err
		}

		requiredCount := good.StoreCount - avaibleCount
		if requiredCount > 0 {
			codesReq = append(codesReq, CodeReq{
				Gtin:     good.Gtin,
				Desc:     good.Desc,
				Required: requiredCount,
			})
		}

	}

	// - Вернуть продукт, описание и недостающее количество кодов
	return codesReq, nil
}

// Добавляет код для печати
func (usecase *UExchange) AddCodeForPrint(
	ctx context.Context,
	code entity.Code,
	source string,
) error {

	return nil
}
