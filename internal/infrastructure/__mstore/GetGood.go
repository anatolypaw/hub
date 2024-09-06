package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Вовзаращет продукт
func (ths *MStore) GetGood(ctx context.Context, gtin string) (entity.Good, error) {
	const op = "mstore.GetGood"

	filter := bson.M{"_id": gtin}
	goods := ths.db.Collection(COLLECTION_GOODS)
	reqResult := goods.FindOne(ctx, filter)

	var good entity.Good
	err := reqResult.Decode(&good)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	if good.Gtin == "" {
		return entity.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	return good, nil
}
