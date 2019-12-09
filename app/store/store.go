package store

// Store ...
type Store interface {
	Job() JobRepository
	JobSetting() JobSettingRepository
	ExchangeJob() ExchangeJobRepository
}
