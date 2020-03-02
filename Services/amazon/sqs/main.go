package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"projects/Services/amazon/sqs/data"
	"projects/Services/amazon/sqs/service"

	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	logfile, err := os.OpenFile("./webapp_sqs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	router.GET("/healthCheck", healthCheck)
	router.POST("/createsqs", createMainQueue)
	router.POST("/sendmessagesqs", sendSQSMessage)

	router.Run(":8090")
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

func healthCheck(ctx *gin.Context) {
	ctx.Status(200)
}

func createMainQueue(ctx *gin.Context) {

	var req data.CreateQueueReq

	ctx.BindJSON(&req)

	log.Println("info: called createMainQueue()")
	log.Println("info: Received parameter", req)

	res := service.CreateQueueService(req)
	ctx.JSON(http.StatusOK, res)

}

func sendSQSMessage(ctx *gin.Context) {

	var req data.SendMessageReq

	ctx.BindJSON(&req)

	log.Println("info: called sendSQSMessage()")
	log.Println("info: Received parameter", req)

	res := service.SendMessageService(req)
	ctx.JSON(http.StatusOK, res)
}
