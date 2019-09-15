package properties

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var prop Config
var config Config

type Config struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	LoginUrl        string `toml:"loginUrl"`
	InfoUrl         string `toml:"infoUrl"`
	EditInfoUrl     string `toml:"editInfoUrl"`
	CheckSessionUrl string `toml:"checkSessionUrl"`
}

func Init() {
	_, err := toml.DecodeFile("../routes/properties/routes-properties.tml", &config)

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
