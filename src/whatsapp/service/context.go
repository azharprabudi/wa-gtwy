package waservice

import (
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/jmoiron/sqlx"
)

// WhatsappService ...
type WhatsappService struct {
	db           *sqlx.DB
	waConn       *whatsapp.Conn
	maxTimeoutQR time.Duration
}
