package main

import (
	"projects/DogOrCat/web/config"
	"projects/DogOrCat/web/controller"
)

func main() {

	controller.StartWebServer(config.Env.Port)

}
