package store

import "github.com/pallid/feelin/app/model"

// Repository ...
type Repository interface {
	Close()
	SaveEntity(*model.QueryResult) error
}

var impl Repository

// SetRepository ...
func SetRepository(repository Repository) {
	impl = repository
}

// Close ...
func Close() {
	impl.Close()
}

// SaveEntity ...
func SaveEntity(q *model.QueryResult) error {
	return impl.SaveEntity(q)
}
