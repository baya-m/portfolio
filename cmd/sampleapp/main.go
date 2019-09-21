package main

import (
	"fmt"
	"git/portfolio/server/config"
	"git/portfolio/server/utils"
)

func main() {
	fmt.Printf("Hello")
	utils.LoggingSetting(config.Config.LogFile)
}
