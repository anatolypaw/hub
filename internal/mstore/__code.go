package mstore

import (
	"context"
	"fmt"
	"hub/internal/entity"

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

// Возвращает код
func (ths *MStore) GetCode(
	ctx context.Context,
	gtin string,
	serial string,
) (entity.FullCode, error) {
	const op = "mongo.GetCode"

	filter := bson.M{"_id": serial}
	codes := ths.db.Collection(gtin)
	reqResult := codes.FindOne(ctx, filter)

	var code entity.FullCode
	err := reqResult.Decode(&code)
	if err != nil {
		return entity.FullCode{}, fmt.Errorf("%s: %w", op, err)
	}
	if code.Serial == "" {
		return entity.FullCode{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	// MAPPING
	mappedCode := entity.FullCode{
		Code: entity.Code{
			Gtin:   gtin,
			Serial: code.Serial,
			Crypto: code.Crypto,
		},
		SourceInfo: code.SourceInfo,
		PrintInfo:  code.PrintInfo,
		Produce:    code.Produce,
		UploadInfo: code.UploadInfo,
	}

	return mappedCode, nil

}
