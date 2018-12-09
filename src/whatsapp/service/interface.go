package waservice

import (
	"github.com/Rhymen/go-whatsapp"
)

// WhatsappServiceInterface ...
type WhatsappServiceInterface interface {
	Login() (*whatsapp.Conn, error)
	Logout(*whatsapp.Conn) error
	SendTextMessagePersonal(*whatsapp.Conn, string, string) error
}
