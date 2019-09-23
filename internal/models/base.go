package models

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
	DbConnection, err := sql.Open("mysql", config.Config.DbProfile)
	if err != nil {
		log.Fatal(err)
	}
	DbConnection.Ping()
	fmt.Println(DbConnection.Ping())
	err = DbConnection.Ping()
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	rows, err := DbConnection.Query("show tables from portfolio")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rows)

}
