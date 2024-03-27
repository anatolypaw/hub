package mstore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// TODO в случае изменения поля printinfo entity.Code,
// может перестать выполняться запрос
// Можно решить полнным маппингом структуры кода
func (ths *MStore) GetCountPrintAvaible(ctx context.Context, gtin string,
) (uint, error) {
	filter := bson.M{"printinfo.avaible": true}
	codes := ths.db.Collection(gtin)
	avaible, err := codes.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return uint(avaible), err
}
