package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/portfolio/internal"

	// "github.com/portfolio/internal/infra"

	_ "github.com/portfolio/internal/infra/persistence"
	config "github.com/portfolio/pkg/configs"
	"github.com/portfolio/pkg/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	fmt.Println("start load db")

	// dbConnection := persistence.DbConnection{}
	// dbConnection.Connect(dbConn)
	// defer dbConn.Close()

	defer fmt.Println("次にDBを閉じます")
	internal.StartWebServer()
	fmt.Println("次がDBConnection")
}
