package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile   string
	Port      int
	DbProfile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("../../init/config.ini")
	if err != nil {
		log.Printf("Cannot load this file :%v", err)
	}
	Config = ConfigList{
		LogFile:   cfg.Section("server").Key("log_file").String(),
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbProfile: cfg.Section("db").Key("profile").String(),
	}

}
