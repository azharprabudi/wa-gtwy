package waservice

import (
	"github.com/Rhymen/go-whatsapp"
)

// WhatsappServiceInterface ...
type WhatsappServiceInterface interface {
	Login() (*whatsapp.Conn, error)
	Logout(*whatsapp.Conn) error
	SetConnectionWhatsapp(*whatsapp.Conn)
	SendTextMessagePersonal(string, string) error
}
