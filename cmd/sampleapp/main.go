package main

import (
	"fmt"

	"github.com/portfolio/internal"
	"github.com/portfolio/internal/infra"
	config "github.com/portfolio/pkg/configs"
	"github.com/portfolio/pkg/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	internal.StartWebServer()
	fmt.Println("次がDBConnection")
	fmt.Println(infra.DbConnection)

}
