package main

import (
	"log"
	"projects/Services/viewer/properties"
	"projects/Services/viewer/viewter"

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
	router.GET("/", viewter.Index)

	//Rooting login,The service will return JSON.
	router.POST("/login", viewter.LoginDS)

	router.Run(":8087")
}
