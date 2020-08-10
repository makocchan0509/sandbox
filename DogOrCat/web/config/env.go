package config

type EnvList struct {
	Port        string
	RabbitUrl   string
	RabbitUser  string
	RabbitPass  string
	RabbitQName string
}

var Env EnvList

func init() {
	/*
		Env = EnvList{
			Port:    os.Getenv("APP_PORT"),
			RabbitUrl: os.Getenv("RABBIT_URL"),
			RabbitUser: os.Getenv("RABBIT_USER"),
			RabbitPass: os.Getenv("RABBIT_PASSWORD"),
			RabbitQName: os.Getenv("RABBIT_QUEUE_NAME")
		}
	*/
	Env = EnvList{
		Port:        "8080",
		RabbitUrl:   "localhost:5672/",
		RabbitUser:  "admin",
		RabbitPass:  "password",
		RabbitQName: "dogOrCatQ",
	}

}
