package mstore

type TerminalParam struct {
	tname string `bson:"_id"`
}

type GoodTerminal struct {
	Gtin string
	Desc string
}

// Возвращает доступные для запрошенного терминала продукты
func (m *MStore) GetGoodsForTerminal(tname string) ([]GoodTerminal, error) {
	return nil, nil
}
