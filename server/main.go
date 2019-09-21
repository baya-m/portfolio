package main

import (
	"config"
	"fmt"
	"server/utils"
)

func main() {
	fmt.Printf("Hello")
	utils.LoggingSetting(config.Config.LogFile)
}
