package store

import (
	"github.com/pallid/feelin/app/model"
)

// Repository ...
type Repository interface {
	Close()
	SaveEntity(*model.QueryResult) error
	GetQueryTextForDeleteData(*model.QueryResult) string
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

// GetQueryTextForDeleteData возвращает текст запроса
// для удаления данных
func GetQueryTextForDeleteData(q *model.QueryResult) string {
	return impl.GetQueryTextForDeleteData(q)
}
