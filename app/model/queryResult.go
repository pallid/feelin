package model

// QueryResult ...
type QueryResult struct {
	BaseModel
	Area           int                      `json:"НомерОбласти"`
	ResultRequest  []map[string]interface{} `json:"РезультатЗапроса"`
	ErrorExecution bool                     `json:"ОшибкаВыполнения"`
	EmptyRequest   bool                     `json:"ПустойЗапрос"`
	ExchangeJobID  string                   `json:"УИД"`
	JobID          string                   `json:"УИД_Пакета"`
	Parameter
}

// Parameter ...
type Parameter struct {
	TableName        string              // имя таблицы
	UseSoftRemove    bool                // использовать мягкое удаление
	SelectFields     []map[string]string // поля отбора
	CompareAllFields bool                // сравнивать по всем полям
	MatchFields      []map[string]string // поля сравнения изменений
}
