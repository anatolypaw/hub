package models

import (
	"time"
)

type Good struct {
	Gtin            string `bson:"_id"`
	Desc            string
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	Upload          bool
	Created         time.Time
}

func (good *Good) ValidateDesc() error {
	if len([]rune(good.Desc)) < 3 || len([]rune(good.Desc)) > 30 {
		return ErrorValidateDesc
	}
	return nil
}

func (good *Good) ValidateGtin() error {
	return ValidateGtin(good.Gtin)
}
