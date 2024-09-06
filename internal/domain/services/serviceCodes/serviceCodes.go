package servicecodes

import (
	"hub/internal/domain/models"
	"time"
)

type ICodesRepo interface {
	Add(code models.FullCode) error
	Get(gtin, serial string) (models.FullCode, error)
	Update(models.FullCode) error

	// Возвращает код, доступный для печати и сразу отмечает его недоступным
	GetForPrint(gtin string) (models.FullCode, error)

	// Атомарный инкремент счетчика
	IncrementCounter(key string) (uint32, error)
}

// Счетчик произведенных кодов
// При старте данные подгружаются из репо
// Вызов Discard декрементирует счетчик произведенных
type ICounters interface {
	// Флаг готовности счетчиков
	Ready() bool

	Produce(gtin, tname, proddate string)
	Produced(gtin, tname, proddate string) uint32

	Discard(gtin, tname, proddate string)
	Discarded(gtin, tname, proddate string) uint32

	PrintAvaibleInc(gtin string)
	PrintAvaibleDec(gtin string)
	PrintAvaible(gtin string) uint32

	PrintID(gtin, tname, proddate string) uint32
	PrintIDNext(gtin, tname, proddate string) uint32
}

type serviceCodes struct {
	repo     ICodesRepo
	counters ICounters
}

func New(repo ICodesRepo) serviceCodes {
	return serviceCodes{
		repo: repo,
	}
}

// Добавить код для печати
func (s *serviceCodes) AddForPrint(gtin, serial, crypto string) error {
	if !s.counters.Ready() {
		return ErrorCountersNotReady
	}

	code, err := models.NewCode(gtin, serial, crypto)
	if err != nil {
		return err
	}

	fullCode := models.FullCode{
		Code:         code,
		LoadedAt:     time.Now(),
		PrintAvaible: true,
		PrintID:      0,
	}

	// Вернет ошибку, если код уже есть
	err = s.repo.Add(fullCode)
	if err != nil {
		return err
	}

	// Увеличиваем счетчик доступных для печати кодов
	s.counters.PrintAvaibleInc(gtin)

	return nil
}

type CodeForPrint struct {
	models.Code
	PrintID uint32
}

// Получить код для печати
func (s *serviceCodes) GetForPrint(gtin, tname, proddate string) (CodeForPrint, error) {
	if !s.counters.Ready() {
		return CodeForPrint{}, ErrorCountersNotReady
	}

	// Валидация входных данных
	if err := models.ValidateGtin(gtin); err != nil {
		return CodeForPrint{}, err
	}

	if err := models.ValidateTName(tname); err != nil {
		return CodeForPrint{}, err
	}

	if err := models.ValidateISODate(proddate); err != nil {
		return CodeForPrint{}, err
	}

	// Поиск свободного для печати кода и установка флага блокировки
	// Свободен тот, где Code.PrintAvaible == true
	fullCode, err := s.repo.GetForPrint(gtin)
	if err != nil {
		return CodeForPrint{}, err
	}

	// Уменьшаем счетчик доступных кодов для печати
	s.counters.PrintAvaibleDec(gtin)

	// Добавляем коду PrintID
	// TODO возможно чтение из разных потоков одного значения!!!
	fullCode.PrintID = s.counters.PrintID(gtin, tname, proddate)

	// И обновляем код в репо
	if err := s.repo.Update(fullCode); err != nil {
		return CodeForPrint{}, err
	}

	// После того, как сделана запись в репо, увеличиваем счетчик
	s.counters.PrintIDNext(gtin, tname, proddate)

	return CodeForPrint{
		Code:    fullCode.Code,
		PrintID: fullCode.PrintID,
	}, nil
}

// Произвести напечатанный код
func (s *serviceCodes) ProducePrinted(gtin, serial, tname, proddate string) error {
	if !s.counters.Ready() {
		return ErrorCountersNotReady
	}

	// Валидация входных данных
	if err := models.ValidateGtin(gtin); err != nil {
		return err
	}

	if err := models.ValidateSerial(serial); err != nil {
		return err
	}

	if err := models.ValidateTName(tname); err != nil {
		return err
	}

	if err := models.ValidateISODate(proddate); err != nil {
		return err
	}

	// Получаем код из репо
	fullCode, err := s.repo.Get(gtin, serial)
	if err != nil {
		return err
	}

	// Обновляем данные о производстве и обновляем в репо
	fullCode.Produced = true
	fullCode.ProdTname = tname
	fullCode.ProdDate = proddate
	fullCode.ProdTime = time.Now()

	// Сохраняем код в репо
	if err := s.repo.Update(fullCode); err != nil {
		return err
	}

	// Увеличиваем счетчик произведенных кодов
	s.counters.Produce(gtin, tname, proddate)

	return nil
}

// Произвести считанный камерой код
func (s *serviceCodes) ProduceReaded(gtin, serial, crypto, tname, proddate string) error {
	if !s.counters.Ready() {
		return ErrorCountersNotReady
	}

	if err := models.ValidateTName(tname); err != nil {
		return err
	}

	if err := models.ValidateISODate(proddate); err != nil {
		return err
	}

	code, err := models.NewCode(gtin, serial, crypto)
	if err != nil {
		return err
	}

	fullCode := models.FullCode{
		Code:     code,
		LoadedAt: time.Now(),

		Produced:  true,
		ProdTime:  time.Now(),
		ProdDate:  proddate,
		ProdTname: tname,
	}

	// Записываем код в репо
	if err := s.repo.Add(fullCode); err != nil {
		return err
	}

	// Увеличиваем счетчик произведенных кодов
	s.counters.Produce(gtin, tname, proddate)

	return nil
}

// Отбраковать код
func (s *serviceCodes) Discard(gtin, serial, tname string) error {
	if !s.counters.Ready() {
		return ErrorCountersNotReady
	}

	if err := models.ValidateGtin(gtin); err != nil {
		return err
	}

	if err := models.ValidateSerial(serial); err != nil {
		return err
	}

	if err := models.ValidateTName(tname); err != nil {
		return err
	}

	// Получить код из репо
	fullCode, err := s.repo.Get(gtin, serial)
	if err != nil {
		return err
	}

	fullCode.Discarded = true
	fullCode.DiscardTname = tname
	fullCode.DiscardTime = time.Now()

	// Обновляем код в репо
	if err := s.repo.Update(fullCode); err != nil {
		return err
	}

	// Увеличиваем счетчик отбракованных кодов
	s.counters.Discard(gtin, tname, fullCode.ProdDate)

	return nil
}

// Получить произведенное и отбракованное кол-во кодов на линии
