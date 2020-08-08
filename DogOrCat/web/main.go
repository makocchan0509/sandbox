package main

import (
	"projects/DogOrCat/web/config"
	"projects/DogOrCat/web/controller"
)

func main() {

	//util.Logging(config.Env.LogPath, config.Config.LogFileName)
	controller.StartWebServer(config.Env.Port)

}
