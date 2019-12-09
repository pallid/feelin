package store

import "github.com/pallid/feelin/app/model"

// JobRepository ...
type JobRepository interface {
	Update(*model.Job) error
	Find(string) (*model.Job, error)
}

// JobSettingRepository ...
type JobSettingRepository interface {
	Update(*model.JobSetting) error
	FindAll(string) (*[]model.JobSetting, error)
}

// ExchangeJobRepository ...
type ExchangeJobRepository interface {
	Update(*model.ExchangeJob) error
	Find(string, string, string) (*model.ExchangeJob, error)
}
