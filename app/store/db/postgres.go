package db

import (
	"database/sql"

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
