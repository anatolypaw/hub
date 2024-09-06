package usecase

import (
	"hub/internal/domain/models"
	"hub/internal/domain/services"
	"log"
)

/*
Предоставляет хранение продуктов и кодов маркировки
Выдает коды для нанесения
*/
type Hub struct {
	sgoods services.Goods
}

func NewHub(sgoods services.Goods) Hub {
	hub := Hub{
		sgoods: sgoods,
	}

	return hub
}

// Добавляет новый продукт
func (hub *Hub) NewGood(good models.Good) error {
	// TODO Проверка прав

	err := hub.sgoods.Add(good)
	log.Println("Добавление нового продукта ", good, "Ошибка: ", err)
	return err
}

// Возвращает продукт по gtin
func (hub *Hub) GetGood(gtin string) (models.Good, error) {
	return hub.sgoods.Get(gtin)
}

// Возвращает все продукты
func (hub *Hub) GetAllGoods() ([]models.Good, error) {
	return hub.sgoods.GetAll()
}
