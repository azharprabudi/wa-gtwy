package wamigration

import "github.com/jmoiron/sqlx"

// DoRunMigration ...
func (m *WhatsappMigration) DoRunMigration() {
	for _, sql := range values {
		m.db.Exec(sql)
	}
}

// NewWhatsappMigration ...
func NewWhatsappMigration(db *sqlx.DB) WhatsappMigrationInterface {
	return &WhatsappMigration{
		db: db,
	}
}
