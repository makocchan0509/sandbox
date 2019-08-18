package main

import (
	"log"
	"projects/Services/routes"
	"projects/Services/routes/properties"

	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	properties.Init()

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	//Rooting login,The service will return JSON.
	router.POST("/login", routes.Login)

	router.Run(":8085")
}
