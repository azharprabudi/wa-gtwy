package waservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/satori/go.uuid"

	"github.com/wa-gtwy/helper/querybuilder/model"

	"github.com/wa-gtwy/helper/querybuilder"

	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/jmoiron/sqlx"

	whatsapp "github.com/Rhymen/go-whatsapp"
	wamodel "github.com/wa-gtwy/src/whatsapp/model"
)

/**
* at this function below, we should check if there
* is existing session, or not. If session is already
* exists, we dont force the user to login again.
* just use previously, but if the session exists
* and is not valid, force user to relogin or
* refresh their session
 */

// Login ...
func (ws *WhatsappService) Login() (*whatsapp.Conn, error) {
	/**
	* open connection to whatsapp first
	 */
	wac, err := whatsapp.NewConn(ws.maxTimeoutQR)
	if err != nil {
		return nil, err
	}

	/**
	* after open connection whatsapp, we should check if
	* the previous session is already exists or not,
	* at different function
	 */
	prevSess, err := ws.readSession()
	if err != nil {
		return nil, err
	}

	/**
	* load previous session, and we should check if
	* that still valid or not
	* when we got the error, we already know if the
	* session is invalid, so we direct user to login
	* again
	 */
	session, err := wac.RestoreSession(*prevSess)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "restoring failed: %v\n", err)
		// return
	} else {
		/**
		* create the barcode at the below, and do login
		* at here
		 */
		qr := make(chan string)

		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()

		session, err = wac.Login(qr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		}
	}

	/**
	* we should save the session
	* to the db, so next time
	* if session exists just use that
	 */
	err = ws.writeSession(&session)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

/**
* get previous session at db
 */
func (ws *WhatsappService) readSession() (*whatsapp.Session, error) {
	/**
	* get last previous session at db, order it by create at desc
	 */
	qbuilder := qb.NewQueryBuilder()

	/**
	* create order
	 */
	order := &qbmodel.Order{
		Key:   "created_at",
		Value: "DESC",
	}

	/**
	* exec query
	 */
	query := qbuilder.QueryWhere("sessions", nil, []*qbmodel.Order{order})

	/**
	* create model from db, and find at db
	* the previous session is exists or not
	 */
	prevSession := new(wamodel.SessionGet)
	err := ws.db.QueryRowx(query).StructScan(&prevSession)
	if err != nil {
		return nil, errors.New("There is no session exists")
	}

	/**
	* if we found the previous session, then we have to
	* parse it into the whatsapp session. But before do it
	* we must to decode that
	 */
	result := new(whatsapp.Session)
	err = json.Unmarshal([]byte(prevSession.SessionCreated.Value), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/**
* set previous version to db
 */
func (ws *WhatsappService) writeSession(session *whatsapp.Session) error {
	/**
	* make it to string before, we insert to
	* tbl session
	 */
	str, err := json.Marshal(*session)
	if err != nil {
		return err
	}
	/**
	* save data to db
	 */
	qbuilder := qb.NewQueryBuilder()
	queryCreate := qbuilder.Create("sessions", &wamodel.SessionCreated{
		ID:        uuid.NewV4(),
		Value:     string(str),
		CreatedAt: time.Now(),
	})
	_, err = ws.db.Exec(queryCreate)
	if err != nil {
		return err
	}
	return nil

}

/**
* the function at below, just focusing on send message to personal user
* specially at the client, we have to sanitize the phone number to be
* like that {country_code}{number}, eg : 63812212121
 */

// SendTextMessagePersonal ...
func (ws *WhatsappService) SendTextMessagePersonal(conn *whatsapp.Conn, destPhoneNumb string, text string) error {
	message := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			FromMe:    true,
			RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", destPhoneNumb),
		},
		Text: text,
	}
	err := conn.Send(message)
	if err != nil {
		return err
	}
	return nil
}

/**
* this function at below, to force the whatsapp web
* to logout from this account
 */

// Logout ...
func (ws *WhatsappService) Logout(wa *whatsapp.Conn) error {
	err := wa.Logout()
	if err != nil {
		return err
	}
	return nil
}

// NewWhatsappService ...
func NewWhatsappService(db *sqlx.DB, maxTimeoutQR time.Duration) WhatsappServiceInterface {
	return &WhatsappService{
		db:           db,
		maxTimeoutQR: maxTimeoutQR,
	}
}
