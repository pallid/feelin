package store

import (
	"github.com/pallid/feelin/app/model"
)

// Repository ...
type Repository interface {
	Close()
	SaveEntity(*model.QueryResult) error
	SetQueryTextForDeleteData(*model.QueryResult)
	SetQueryTextForSelectData(*model.QueryResult)
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

// SetQueryTextForDeleteData возвращает текст запроса
// для удаления данных
func SetQueryTextForDeleteData(q *model.QueryResult) {
	impl.SetQueryTextForDeleteData(q)
}

// SetQueryTextForSelectData возвращает текст запроса
// для удаления данных
func SetQueryTextForSelectData(q *model.QueryResult) {
	impl.SetQueryTextForSelectData(q)
}
