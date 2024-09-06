package models

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrorValidateGtin    = errors.New("некорректный формат GTIN")
	ErrorValidateSerial  = errors.New("некорректный формат serial")
	ErrorValidateCrypto  = errors.New("некорректный формат crypto")
	ErrorValidateDesc    = errors.New("некорректное название")
	ErrorValidateISODate = errors.New("некорректный формат даты. Используйте ISO 8601")
	ErrorValidateTName   = errors.New("некорректный формат имени линии")
)

var (
	gtinValidator   = regexp.MustCompile(`^0\d{13}$`)
	serialValidator = regexp.MustCompile(`^([a-zA-Z0-9]|[!"%&'*+\-.\/_,:;=<>?()]){6}$`)
	cryptoValidator = regexp.MustCompile(`^([a-zA-Z0-9]|[!"%&'*+\-.\/_,:;=<>?()]){4}$`)
)

func ValidateGtin(gtin string) error {
	if !gtinValidator.MatchString(string(gtin)) {
		return ErrorValidateGtin
	}
	return nil
}

func ValidateSerial(serial string) error {
	if !serialValidator.MatchString(serial) {
		return ErrorValidateSerial
	}
	return nil
}

func ValidateCrypto(crypto string) error {
	if !cryptoValidator.MatchString(crypto) {
		return ErrorValidateCrypto
	}
	return nil
}

func ValidateTName(tname string) error {
	if len(tname) == 0 {
		return ErrorValidateTName
	}
	return nil
}

func ValidateISODate(date string) error {
	// Определяем формат для разбора ISO 8601 даты
	layout := "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		return ErrorValidateISODate
	}
	return nil
}
