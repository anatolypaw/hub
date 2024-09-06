package mock

import (
	"errors"
	"hub/internal/domain/models"
)

type repoGoods struct {
	store map[string]models.Good
}

func NewRepoGoods() repoGoods {
	rp := repoGoods{
		store: make(map[string]models.Good),
	}
	return rp
}

func (rp *repoGoods) Add(good models.Good) error {
	rp.store[good.Gtin] = good
	return nil
}

func (rp *repoGoods) Get(gtin string) (models.Good, error) {
	good, ok := rp.store[gtin]
	if !ok {
		return models.Good{}, errors.New("продукт не найден")
	}
	return good, nil
}

func (rp *repoGoods) GetAll() ([]models.Good, error) {
	goods := make([]models.Good, 0, len(rp.store))

	for _, good := range rp.store {
		goods = append(goods, good)
	}
	return goods, nil
}
