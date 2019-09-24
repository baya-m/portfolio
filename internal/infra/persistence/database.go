package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/portfolio/pkg/configs"
)

var DbConnection *sql.DB

func init() {
	fmt.Println("start load db")
	var err error
	DbConnection, err = sql.Open("mysql", config.Config.DbProfile)
	if err != nil {
		log.Fatal(err)
	}
	if err := DbConnection.Ping(); err != nil {
		log.Printf("failed to ping by error '%#v'", err)
	}
}
