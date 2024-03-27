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
	Serial    string `bson:"_id"`
	Crypto    string
	PrintInfo PrintInfo
	ProdInfo  []ProdInfo
}

// Информация выгрузки
type ProdInfo struct {
	Time time.Time // Время события

	Type     string    // Тип события. discard / produce
	ProdDate time.Time // Дата производства, если тип produce
	Tname    string    // Имя линии, на которой произведено или отбраковано

	Uploaded   bool
	UploadTime time.Time // вермя выгрузки
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
