package wamigration

import "github.com/jmoiron/sqlx"

// WhatsappMigration ...
type WhatsappMigration struct {
	db *sqlx.DB
}
