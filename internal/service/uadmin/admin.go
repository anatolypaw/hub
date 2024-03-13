package uadmin

import (
	"context"
	"hub/internal/entity"
	"time"
)

/*
Методы для управления хранением марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/
type iGoodRepo interface {
	Add(context.Context, entity.Good) error
	GetGood(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type UAdmin struct {
	goodRepo iGoodRepo
}

func New(goodRepo iGoodRepo) UAdmin {
	return UAdmin{
		goodRepo: goodRepo,
	}
}

// Добавляет новый продукт
// Валидация gtin, desc
// Ошибка, если такой продукт с таким gtin уже существует
func (u *UAdmin) AddGood(ctx context.Context, good entity.Good,
) error {
	err := good.ValidateGtin()
	if err != nil {
		return err
	}

	err = good.ValidateDesc()
	if err != nil {
		return err
	}

	good.CreatedAt = time.Now()
	return u.goodRepo.Add(ctx, good)
}

func (ths *UAdmin) GetAllGoods(ctx context.Context,
) ([]entity.Good, error) {
	// TODO валидировать ответ хранилища
	// на корректность gtin
	return ths.goodRepo.GetAll(ctx)
}
