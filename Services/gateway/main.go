package main

import (
	"log"
	"projects/Services/routes"
	"projects/Services/routes/properties"

	_ "net/http/pprof"

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

	router.OPTIONS("/:uri", routes.Options)

	router.POST("/login", routes.Login)
	router.POST("/getInfo", routes.GetInfoLists)
	router.POST("/editInfo", routes.EditInfo)
	router.POST("/checkSession", routes.CheckSession)

	router.Run(":8085")
}
