package main

import (
	"fmt"

	config "github.com/portfolio/pkg/configs"

	"github.com/portfolio/pkg/utils"
)

func main() {
	fmt.Printf("%v", config.Config.Port)
	utils.LoggingSetting(config.Config.LogFile)
}
