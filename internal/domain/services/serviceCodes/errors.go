package servicecodes

import "errors"

var (
	ErrorCountersNotReady = errors.New("кэш не готов, идет загрузка даннных")
)
