package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ...
	"github.com/pallid/feelin/app/model"
)

// PostgresRepository ...
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgres ...
func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

// Close ...
func (r *PostgresRepository) Close() {
	r.db.Close()
}

// SaveEntity ...
func (r *PostgresRepository) SaveEntity(q *model.QueryResult) error {
	return nil
}

// SetQueryTextForDeleteData возвращает текст запроса
// для удаления данных
func (r *PostgresRepository) SetQueryTextForDeleteData(q *model.QueryResult) {
	var t string
	switch {
	case q.HardRemoval:
		t = `DELETE from %s WHERE area = %d`
	default:
		t = `UPDATE %s SET deleted_at = NULL area = %d`
	}
	q.DeleteData = fmt.Sprintf(t, q.TableName, q.Area)
}
