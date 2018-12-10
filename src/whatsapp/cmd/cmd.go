package wa

import (
	"time"

	"github.com/wa-gtwy/src/whatsapp/db"
	"github.com/wa-gtwy/src/whatsapp/migration"
	"github.com/wa-gtwy/src/whatsapp/service"
)

// InitializeWA ...
func InitializeWA() (waservice.WhatsappServiceInterface, error) {

	/**
	* open db connection for
	* whatsapp only
	 */
	db := wadb.NewWhatsappDB()
	dbConn, err := db.OpenDBConnection()
	if err != nil {
		return nil, err
	}

	/**
	* run migration db
	 */
	migrate := wamigration.NewWhatsappMigration(dbConn)
	migrate.DoRunMigration()

	/**
	* create whatsapp service
	 */
	service := waservice.NewWhatsappService(dbConn, 100000*time.Second)
	waconn, err := service.Login()
	if err != nil {
		return nil, err
	}

	/**
	* after we receive the connection whatsapp,
	* we should set to our service
	 */

	service.SetConnectionWhatsapp(waconn)
	return service, nil
}
