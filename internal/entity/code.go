package entity

import (
	"errors"
	"regexp"
	"time"
)

type Code struct {
	Gtin   string
	Serial string
	Crypto string
}

type FullCode struct {
	Serial     string `bson:"_id"`
	Crypto     string
	Type       string         // print || read
	PrintInfo  PrintInfo      `bson:",omitempty"`
	ProduceLog []ProducedInfo `bson:",omitempty"`

	ProdDate string // Дата производства
	Tname    string // Имя линии, на которой произведено
	Discard  bool   // Брак
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Tname        string    // Имя линии фасовки, где он был нанесен или считан камерой
	Time         time.Time // Время, когда код был нанесен или считан на линии
	ProdDate     string    // Дата производства продукта 2023-10-09
	Discard      bool      // True - операция отбраковки кода
	UploadTime   time.Time // Информация о выгрузке во внешнюю систему
	UploadStatus string    // Информация о выгрузке во внешнюю систему
}

// Информация, связанная с печатью
type PrintInfo struct {
	Sname    string    // Откуда загружен, например с сервера "server main"
	Loaded   time.Time // Время получения кода
	Avaible  bool      // Флаг, что код доступен для печати
	Uploaded time.Time // Время выдачи кода из базы
	Tname    string    // Имя линии, куда передан код

	PrintID uint32 // Последовательный номер кода в партии и линии, присваивается при выдаче кода на печать
}

func ValidateSerial(serial string) error {
	r := regexp.MustCompile(`^([a-zA-Z0-9]|[!"%&'*+\-.\/_,:;=<>?]){6}$`)

	if !r.MatchString(serial) {
		return errors.New("некорректный формат serial")
	}
	return nil
}

func ValidateCrypto(crypto string) error {
	r := regexp.MustCompile(`^([a-zA-Z0-9]|[!"%&'*+\-.\/_,:;=<>?]){4}$`)

	if !r.MatchString(crypto) {
		return errors.New("некорректный формат crypto")
	}

	return nil
}
