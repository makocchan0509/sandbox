package properties

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var prop Config
var config Config

type Config struct {
	Mysql   MysqlConfig
	Redis   RedisConfig
	Session SessionConfig
}

type MysqlConfig struct {
	MysqlHost   string `toml:"mysqlhost"`
	MysqlPort   string `toml:"mysqlport"`
	MysqlUser   string `toml:"mysqluser"`
	MysqlPass   string `toml:"mysqlpass"`
	MysqlDBName string `toml:"mysqldbname"`
}

type RedisConfig struct {
	RedisHost string `toml:"redishost"`
	RedisPort string `toml:"redisport"`
	RedisPrtc string `toml:"redisprotocol"`
}

type SessionConfig struct {
	Url string `toml:"url"`
}

func Init() {

	_, err := toml.DecodeFile("../common/properties/common-properties.tml", &config)

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
