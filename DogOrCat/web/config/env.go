package config

type EnvList struct {
	Port      string
	RedisHost string
	RedisPort string
}

var Env EnvList

func init() {
	/*
		Env = EnvList{
			Port:    os.Getenv("APP_PORT"),
			RedisHost:	os.Getenv("REDIS_HOST")
			RedisPort:	os.Getenv("REDIS_PORT")
		}
	*/
	Env = EnvList{
		Port:      "8080",
		RedisHost: "127.0.0.1",
		RedisPort: "6379",
	}

}
