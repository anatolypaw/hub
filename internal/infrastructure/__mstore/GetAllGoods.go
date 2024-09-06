package mstore

import (
	"context"
	"hub/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MStore) GetAllGoods(ctx context.Context) ([]entity.Good, error) {

	filter := bson.M{}
	cursor, err := m.db.Collection(COLLECTION_GOODS).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var goods []entity.Good

	// Проходим по всем документам и добавляем их в список.
	for cursor.Next(ctx) {
		var good entity.Good

		if err := cursor.Decode(&good); err != nil {
			return nil, err
		}
		goods = append(goods, good)
	}
	return goods, nil
}
