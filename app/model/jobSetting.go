package model

// JobSetting ...
type JobSetting struct {
	BaseModel
	JobID  string `json:"ИдентификаторЗадания"`
	Option string
	Value  string
}
