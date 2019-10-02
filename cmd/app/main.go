package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/portfolio/internal"

	_ "github.com/portfolio/internal/infra/persistence"
	"github.com/portfolio/config"
	"github.com/portfolio/pkg/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	internal.StartWebServer()
}
