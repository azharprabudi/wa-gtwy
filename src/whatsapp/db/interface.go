package wadb

import (
	"github.com/jmoiron/sqlx"
)

// WhatsappDBInterface ...
type WhatsappDBInterface interface {
	OpenDBConnection() (*sqlx.DB, error)
}
