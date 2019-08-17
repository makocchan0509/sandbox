package properties

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var prop Config
var config Config

type Config struct {
	Login LoginConfig
}

type LoginConfig struct {
	Url string `toml:"url"`
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
