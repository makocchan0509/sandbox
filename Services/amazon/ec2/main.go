package main

import (
	"io"
	"log"
	"os"

	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	logfile, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()

	colog.SetOutput(io.MultiWriter(logfile, os.Stdout))
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	router.OPTIONS("/:uri", options)
	router.GET("/outputInfo", outputInfo)
	router.GET("/outputError", outputError)

	router.Run(":8080")
}

//Handling OPTIONS method
func options(ctx *gin.Context) {

	log.Println("info: Called OPTONS method.")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Max-Age", "86400")

	// ブラウザからリクエストヘッダーの送信を許可
	ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Status(200)
}

func outputInfo(ctx *gin.Context) {
	name, err := os.Hostname()
	if err != nil {
		log.Println("error:", err.Error())
	}
	log.Println("info: called outputInfo hostname>-- ", name)
	ctx.Status(200)
}

func outputError(ctx *gin.Context) {
	name, err := os.Hostname()
	if err != nil {
		log.Println("error:", err.Error())
	}
	log.Println("error: called outputError hostname>-- ", name)
	ctx.Status(200)
}

func example(x int, y int) (z int) {

	z = x + y

	return z
}
