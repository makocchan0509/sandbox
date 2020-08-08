package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFileName string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("fail load file.")
		os.Exit(1)
	}
	Config = ConfigList{
		LogFileName: cfg.Section("logging").Key("logFileName").String(),
	}
}
