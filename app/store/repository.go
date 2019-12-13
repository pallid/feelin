package store

import "github.com/pallid/feelin/app/model"

type EntityRepository interface {
	Close()
	Save(*model.QueryResult) error
}
