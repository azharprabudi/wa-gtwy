package waservice

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// WhatsappService ...
type WhatsappService struct {
	db           *sqlx.DB
	maxTimeoutQR time.Duration
}
