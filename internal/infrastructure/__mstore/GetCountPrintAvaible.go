package mstore

import (
	"context"
	"hub/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает количество доступных кодов для печати
func (m *MStore) GetCountPrintAvaible(ctx context.Context, gtin string) (int64, error) {
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"printinfo.avaible": true}
	count, err := m.db.Collection(gtin).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
