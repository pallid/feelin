package pgstore

import (
	"github.com/pallid/feelin/app/model"
)

// EntityRepository ...
type EntityRepository struct {
	store *Store
}

// Close ...
func (r *EntityRepository) Close() {
	r.store.db.Close()
}

// Save ...
func (r *EntityRepository) Save(u *model.QueryResult) error {

	// return r.store.db.QueryRow(
	// 	"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
	// 	u.Email,
	// 	u.EncryptedPassword,
	// ).Scan(&u.ID)

	return nil
}
