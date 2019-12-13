package store

import "github.com/pallid/feelin/app/model"

// JobRepository ...
type JobRepository interface {
	Close()
	Update(*model.Job) error
	Find(string) (*model.Job, error)
}

// JobSettingRepository ...
type JobSettingRepository interface {
	Close()
	Update(*model.JobSetting) error
	FindAll(string) (*[]model.JobSetting, error)
}

// ExchangeJobRepository ...
type ExchangeJobRepository interface {
	Close()
	Update(*model.ExchangeJob) error
	Find(string, string, string) (*model.ExchangeJob, error)
}
