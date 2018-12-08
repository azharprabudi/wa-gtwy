package model

import (
	"time"
)

type ReadSessionFromDB struct {
	ID        uuid.UUID  `db:"id"`
	Value     string     `db:"value"`
	CreatedAt time.Time  `db:"created_at"`
	ExpiredAt *time.Time `db:"expired_at"`
}
