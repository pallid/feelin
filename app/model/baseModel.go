package model

import "time"

// BaseModel ...
type BaseModel struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
