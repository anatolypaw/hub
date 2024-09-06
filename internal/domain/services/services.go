package services

import "hub/internal/domain/models"

// Сервис работы с продуктами
type Goods interface {
	Add(good models.Good) error
	Get(gtin string) (models.Good, error)
	GetAll() ([]models.Good, error)
}

// Сервис работы с кодами для маркировки
type Codes interface {
}

type Services struct {
	Goods
	Codes
}
