package entity

import (
	"fmt"
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

func (ths *Good) ValidateDesc() error {
	if len([]rune(ths.Desc)) < 3 || len([]rune(ths.Desc)) > 30 {
		return fmt.Errorf("требуется длина 3 < %d < 30 символов",
			len([]rune(ths.Desc)))
	}
	return nil
}

func (ths *Good) ValidateGtin() error {
	return ValidateGtin(ths.Gtin)
}
