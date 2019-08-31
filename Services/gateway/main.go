package main

import (
	"log"
	"projects/Services/routes"
	"projects/Services/routes/properties"

	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
)

func main() {

	// 妙子大好き！
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

	router.Run(":8085")
}
