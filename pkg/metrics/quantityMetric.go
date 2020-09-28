package metrics

import "time"

// QuantityMetric ...
type QuantityMetric struct {
	DateMetric time.Time // Дата метрики
	Area       string    // Область
	TableName  string    // Имя таблицы
	Value      int       // Значение метрики
	Hash       string    // Строка хеш суммы
}
