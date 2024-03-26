package mstore

import "context"

// Выводит количество фасованных кодов дя запрошенной линии, продукта, даты
func (m *MStore) GetProducedCount(ctx context.Context, tname string, gtin string, date string) (int, error) {

	return 0, nil
}
