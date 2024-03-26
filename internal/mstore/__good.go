package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (ths *MStore) GetAll(ctx context.Context) ([]entity.Good, error) {
	const op = "mongo.GetAll"

	filter := bson.M{}
	goods := ths.db.Collection(COLLECTION_GOODS)
	cursor, err := goods.Find(ctx, filter)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	goods_dto := []Good_dto{}
	err = cursor.All(context.TODO(), &goods_dto)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// MAPPING
	mappedGoods := []entity.Good{}
	for _, good_dto := range goods_dto {
		mappedGoods = append(mappedGoods, entity.Good{
			Gtin:            good_dto.Gtin,
			Desc:            good_dto.Desc,
			StoreCount:      good_dto.StoreCount,
			GetCodeForPrint: good_dto.GetCodeForPrint,
			AllowProduce:    good_dto.AllowProduce,
			AllowPrint:      good_dto.AllowPrint,
			Upload:          good_dto.Upload,
			CreatedAt:       good_dto.CreatedAt,
		})
	}

	return mappedGoods, nil
}
