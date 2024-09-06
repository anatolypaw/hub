package app

import (
	"hub/internal/adapters/rest"
	"hub/internal/domain/models"
	sgoods "hub/internal/domain/services/serviceGoods"
	"hub/internal/domain/usecase"
	"hub/internal/infrastructure/mock"
	"log"
)

func Run() {
	// Репозитории
	repoGoods := mock.NewRepoGoods()

	// Сервисы
	sgoods := sgoods.New(&repoGoods)

	// Юзкейсы
	hub := usecase.NewHub(&sgoods)

	hub.NewGood(models.Good{
		Gtin: "01111111111111",
		Desc: "ass",
	})

	// Адаптеры
	rest := rest.New(&hub)
	go log.Print(rest.Run(":80"))

}
