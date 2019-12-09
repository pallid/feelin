package model

// Job ...
type Job struct {
	BaseModel
	JobID  string `json:"ИдентификаторЗадания"`
	Status string `json:"Состояние"`
	Priod  string `json:"Дата"`
}
