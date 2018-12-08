package waservice

import (
	"time"

	"github.com/jmoiron/sqlx"

	whatsapp "github.com/Rhymen/go-whatsapp"
)

/*
* at this function below, we should check if there
* is existing session, or not. If session is already
* exists, we dont force the user to login again.
* just use previously, but if the session exists
* and is not valid, force user to relogin or
* refresh their session
 */
func (ws *WhatsappService) Login() (*whatsapp.Conn, error) {
	/**
	* open connection to whatsapp first
	 */
	wac, err := whatsapp.NewConn(ws.MaxTimeoutQR)
	if err != nil {
		return nil, err
	}

	/**
	* after open connection whatsapp, we should check if
	* the previous session is already exists or not,
	* at different function
	 */
	ws.readSession()
}

/**
* get previous session at db
 */
func (ws *WhatsappService) readSession() {

}

// SendTextMessage ...
func (ws *WhatsappService) SendTextMessage(destPhoneNumb string, text string) {

}

// NewWhatsappService ...
func NewWhatsappService(db *sqlx.DB, maxTimeoutQR time.Duration) WhatsappServiceInterface {
	return &WhatsappService{
		db:           db,
		maxTimeoutQR: maxTimeoutQR,
	}
}
