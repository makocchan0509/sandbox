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
	GatewayUrl string `toml:"gatewayUrl"`
}

func Init() {
	_, err := toml.DecodeFile("./properties/viewter-properties.tml", &config)

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
