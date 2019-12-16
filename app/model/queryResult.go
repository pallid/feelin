package model

import "github.com/pallid/feelin/pkg/comparekv"

// QueryResult основная модель приложения
type QueryResult struct {
	Area           int                      `json:"НомерОбласти"`
	ResultRequest  []map[string]interface{} `json:"РезультатЗапроса"`
	ErrorExecution bool                     `json:"ОшибкаВыполнения"`
	EmptyRequest   bool                     `json:"ПустойЗапрос"`
	ExchangeJobID  string                   `json:"УИД"`
	JobID          string                   `json:"УИД_Пакета"`
	Options
	RequestTexts
}

// Options описание настроек
// для обработки результата запроса
type Options struct {
	Description      string   `json:"НаименованиеЗадания"`
	TableName        string   `json:"ИмяТаблицы"`
	HardRemoval      bool     `json:"ПолноеУдаление"`
	SelectionFields  []string `json:"ПоляОтбора"`
	ComparionFields  []string `json:"ПоляСравнения"`
	CompareAllFields bool     `json:"СравниватьПоВсемПолям"`
}

// RequestTexts тексты запросов
type RequestTexts struct {
	DeleteData string
	SelectData string
}

// UseHardDelete возвращает признак полного удаления
func (q *QueryResult) UseHardDelete() bool {
	return q.HardRemoval
}

// NeedCompareAllFields возвращает признак того,
// что необходимо сравнивать все поля
func (q *QueryResult) NeedCompareAllFields() bool {
	return q.CompareAllFields
}

// GetTargetTableName возвращает имя таблицы,
// над которой будут выполняться действия
func (q *QueryResult) GetTargetTableName() string {
	return q.TableName
}

// GetSelectionFields возвращает поля таблицы,
// по которым необходимо выполнять отбор
func (q *QueryResult) GetSelectionFields() []string {
	return q.SelectionFields
}

// GetComparisonFields возвращает поля таблицы,
// по которым необходимо выполнять сравнение
func (q *QueryResult) GetComparisonFields() []string {
	return q.ComparionFields
}

// CompareFields метод сравнения структур
// возвращает признак равенства структур
func (q *QueryResult) CompareFields(struct1, struct2 map[string]interface{}) (areEqual bool) {
	if q.HardRemoval {
		return true
	}
	// Сравниваем по всем полям первой структуры
	switch {
	case q.CompareAllFields:

	default:
		astr := q.GetComparisonFields()
		areEqual = comparekv.CompareFields(struct1, struct2, astr...)
	}
	return areEqual
}
