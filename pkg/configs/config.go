package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string
	Port    int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Cannot load this file :%v", err)
	}
	Config = ConfigList{
		LogFile: cfg.Section("server").Key("log_file").String(),
		Port:    cfg.Section("web").Key("port").MustInt(),
	}
	fmt.Printf("%v", Config.Port)
	fmt.Printf("%v", Config.LogFile)
}
