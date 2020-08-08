package config

type EnvList struct {
	LogPath string
	Port    string
}

var Env EnvList

func init() {
	/*
		Env = EnvList{
			LogPath: os.Getenv("APP_LOGPATH"),
			Port:    os.Getenv("APP_PORT"),
		}
	*/
	Env = EnvList{
		//LogPath: ".",
		Port: "8080",
	}

}
