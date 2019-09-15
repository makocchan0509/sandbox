package sqls

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var prop Config
var config Config

type Config struct {
	Sql SqlConfig
}

type SqlConfig struct {
	SelectInfoList string `toml:"Select_InfomationList"`
	InsertInfo     string `toml:"Insert_Information"`
	UpdateInfo     string `toml:"Update_Information"`
	DeleteInfo     string `toml:"Delete_Information"`
}

func Init() {

	_, err := toml.DecodeFile("./sqls/informartion-sqls-properties.tml", &config)

	if err != nil {
		fmt.Println(err.Error())
	}
	setProp(config)
}

func setProp(config Config) {
	prop = config
}

func GetProp() Config {
	return prop
}
