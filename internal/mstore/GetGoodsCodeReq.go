package mstore

import (
	"context"
	"hub/internal/entity"
)

type CodeReq struct {
	Gtin     string
	Desc     string
	Required int64
}

func (m *MStore) GetGoodsCodeReq(ctx context.Context) ([]CodeReq, error) {
	// - Получить продукты
	goods, err := m.GetAllGoods(ctx)
	if err != nil {
		return nil, err
	}

	// - Выбрать те, для которых включено наполнения кодами
	var goodsAvaibleForPrint []entity.Good
	for _, good := range goods {
		if good.GetCodeForPrint {
			goodsAvaibleForPrint = append(goodsAvaibleForPrint, good)
		}
	}

	// - Для каждого продукта получить доступное количество кодов
	var codesReq []CodeReq
	for _, good := range goodsAvaibleForPrint {
		avaibleCount, err := m.GetCountPrintAvaible(ctx, good.Gtin)
		if err != nil {
			return nil, err
		}

		requiredCount := int64(good.StoreCount) - avaibleCount
		codesReq = append(codesReq, CodeReq{
			Gtin:     good.Gtin,
			Desc:     good.Desc,
			Required: requiredCount,
		})

	}

	// - Вернуть продукт, описание и недостающее количество кодов
	return codesReq, nil
}
