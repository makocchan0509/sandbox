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
	router.LoadHTMLGlob("../static/views/*.html")
	router.Static("/static", "../static")

	properties.Init()

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	//Rooting index page.
	router.GET("/", routes.Index)

	//Rooting login,The service will return JSON.
	router.POST("/login", routes.Login)

	//Rooting login and display page.
	router.POST("/loginDS", routes.LoginDS)

	router.Run(":8085")
}
