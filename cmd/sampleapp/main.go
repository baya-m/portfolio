package main

import (
	"github.com/portfolio/internal"
	config "github.com/portfolio/pkg/configs"
	"github.com/portfolio/pkg/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	internal.StartWebServer()
}
