package persistence

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/portfolio/pkg/config"
)

var DbConnection *sql.DB

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
	return nil
}

func init() {
	log.Println("start load db")
	var err error
	DbConnection, err = sql.Open("mysql", config.Config.DbProfile)
	if err != nil {
		log.Fatal(err)
	}
	if err := DbConnection.Ping(); err != nil {
		log.Fatalf("failed to ping by error '%#v'", err)
	}
}
