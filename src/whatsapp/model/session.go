package wamodel

import (
	"time"

	"github.com/satori/go.uuid"
)

// SessionGet ...
type SessionGet struct {
	SessionCreated
}

// SessionCreated ...
type SessionCreated struct {
	ID        uuid.UUID `db:"id"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
}
