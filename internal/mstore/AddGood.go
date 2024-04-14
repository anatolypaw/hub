package mstore

import (
	"context"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// возвращает все поля добавленного продукта
func (m *MStore) AddGood(ctx context.Context, sname string, gtin string, desc string) error {
	// Логгирование
	const op = "mstore.AddGood"
	logger := m.logger.With("func", op).
		With("sname", sname).
		With("gtin", gtin)

	var err error

	start := time.Now()
	defer func() {
		since := time.Since(start)
		logger = logger.With("err", err, "duration", since)
		if err != nil {
			logger.Warn("Response")
		} else {
			logger.Info("Response")
		}
	}()

	// Добавляем продукт в бд
	err = entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	good := entity.Good{
		Gtin:            gtin,
		Desc:            desc,
		GetCodeForPrint: true,
		AllowProduce:    true,
		Created:         time.Now(),
	}

	collect := m.db.Collection(COLLECTION_GOODS)
	_, err = collect.InsertOne(ctx, good)
	if err != nil {
		return err
	}

	// создаем индекс для коллекции кодов по значению printinfo.avaible
	coll := m.db.Collection(gtin)
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "printinfo.avaible", Value: 1}},
	}

	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	// создаем индекс для коллекции кодов по значению prodinfo.uploaded
	indexModel2 := mongo.IndexModel{
		Keys: bson.D{{Key: "prodinfo.uploaded", Value: 1}},
	}

	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel2)
	if err != nil {
		return err
	}

	// создаем индекс для коллекции кодов по значению printinfo.printid
	indexModel3 := mongo.IndexModel{
		Keys: bson.D{{Key: "printinfo.printid", Value: 1}},
	}

	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel3)
	if err != nil {
		return err
	}

	return nil
}
