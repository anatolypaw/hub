package models

import (
	"time"
)

func NewCode(gtin, serial, crypto string) (Code, error) {
	if err := ValidateGtin(gtin); err != nil {
		return Code{}, err
	}

	if err := ValidateSerial(serial); err != nil {
		return Code{}, err
	}

	if err := ValidateCrypto(crypto); err != nil {
		return Code{}, err
	}

	return Code{
		Gtin:   gtin,
		Serial: serial,
		Crypto: crypto,
	}, nil
}

type Code struct {
	Gtin   string
	Serial string
	Crypto string
}

type FullCode struct {
	Code
	LoadedAt     time.Time // Время получения кода
	PrintID      uint32    // Последовательный номер кода в партии и линии, присваивается при выдаче кода на печать
	PrintAvaible bool      // Флаг доступности кода для печати

	Produced       bool      // Произведен
	ProdTime       time.Time // Время производства
	ProdDate       string    // Дата производства 2024-02-02
	ProdTname      string    // Имя линии, на которой произведено
	ProdUploaded   bool      // Флаг, что состояние производства выгружено во внешнюю систему
	ProdUploadTime time.Time // Время выгрузки

	Discarded         bool      // Отбракован
	DiscardTime       time.Time // Время события
	DiscardTname      string    // Имя линии, на которой отбраковано
	DiscardUploaded   bool      // Флаг, что состояние отбраковки выгружено во внешнюю систему
	DiscardUploadTime time.Time // Время выгрузки
}
