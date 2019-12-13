package pgstore

import (
	"database/sql"

	"github.com/pallid/feelin/app/store"
)

// Store ...
type Store struct {
	db               *sql.DB
	entityRepository *EntityRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Close ...
func (r *Store) Close() {
	r.db.Close()
}

// Message ...
func (r *Store) Message() store.EntityRepository {
	if r.entityRepository != nil {
		return r.entityRepository
	}

	r.entityRepository = &EntityRepository{
		store: r,
	}

	return r.entityRepository
}
