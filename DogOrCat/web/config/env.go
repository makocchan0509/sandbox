package config

type EnvList struct {
	Port        string
	RabbitUrl   string
	RabbitUser  string
	RabbitPass  string
	RabbitQName string
	MysqlHost   string
	MysqlPort   string
	MysqlUser   string
	MysqlPass   string
	MysqlDBName string
}

var Env EnvList

func init() {
	/*
		Env = EnvList{
			Port:        os.Getenv("APP_PORT"),
			RabbitUrl:   os.Getenv("RABBIT_URL"),
			RabbitUser:  os.Getenv("RABBIT_USER"),
			RabbitPass:  os.Getenv("RABBIT_PASSWORD"),
			RabbitQName: os.Getenv("RABBIT_QUEUE_NAME"),
			MysqlHost:   os.Getenv("MYSQL_HOST"),
			MysqlPort:   os.Getenv("MYSQL_PORT"),
			MysqlUser:   os.Getenv("MYSQL_USER"),
			MysqlPass:   os.Getenv("MYSQL_PASSWORD"),
			MysqlDBName: os.Getenv("MYSQL_DBNAME"),
		}
	*/
	Env = EnvList{
		Port:        "8080",
		RabbitUrl:   "localhost:5672/",
		RabbitUser:  "admin",
		RabbitPass:  "password",
		RabbitQName: "dogOrCatQ",
		MysqlHost:   "localhost",
		MysqlPort:   "3306",
		MysqlUser:   "admin",
		MysqlPass:   "password",
		MysqlDBName: "dogorcat",
	}
}
