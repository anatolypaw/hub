package sgoods

import (
	"hub/internal/domain/models"
	"time"
)

/*
Сервис хранит продукты и их параметры.
Валидирует данные, при добавлении нового продукта.
*/
type ServiceGoods struct {
	repo IGoodsRepo
}

func New(repo IGoodsRepo) ServiceGoods {
	return ServiceGoods{repo: repo}
}

// Интерфейс репозитория
type IGoodsRepo interface {
	Add(good models.Good) error
	Get(gtin string) (models.Good, error)
	GetAll() ([]models.Good, error)
}

// Валидирует и добавляет новый продукт в хранилище
func (s *ServiceGoods) Add(good models.Good) error {
	if err := good.ValidateDesc(); err != nil {
		return err
	}

	if err := good.ValidateGtin(); err != nil {
		return err
	}

	good.Created = time.Now()
	return s.repo.Add(good)
}

// Возвращает продукт по его GTIN
func (s *ServiceGoods) Get(gtin string) (models.Good, error) {
	if err := models.ValidateGtin(gtin); err != nil {
		return models.Good{}, err
	}

	return s.repo.Get(gtin)
}

// Возвращает все продукты
func (s *ServiceGoods) GetAll() ([]models.Good, error) {
	return s.repo.GetAll()
}
