package wadb

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// OpenDBConnection ...
func (w *WhatsappDB) OpenDBConnection() (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", w.getEnv("WA_DB_HOST"), w.getEnv("WA_DB_PORT"), w.getEnv("WA_DB_USER"), w.getEnv("WA_DB_PASS"), w.getEnv("WA_DB_NAME"))

	// connect db
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (w *WhatsappDB) getEnv(key string) string {
	str := os.Getenv(key)
	return str
}

// NewWhatsappDB ...
func NewWhatsappDB() WhatsappDBInterface {
	return &WhatsappDB{}
}
