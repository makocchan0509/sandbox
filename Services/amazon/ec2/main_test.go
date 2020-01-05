package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/comail/colog"
)

func TestExample(t *testing.T) {

	x := 3
	y := 5
	correct := 8

	z := example(x, y)

	if correct != z {
		t.Fatal("Failed test..")
	}
}

// func TestOptions(t *testing.T) {

// 	//init log file.
// 	initlogging()

// 	req, _ := http.NewRequest("GET", "/outputInfo", nil)
// 	param := gin.Param{"classificationId", "1"}
// 	params := gin.Params{param}

// 	var context *gin.Context
// 	context = &gin.Context{Request: req, Params: params}

// 	options(context)

// 	val, _ := context.Get("Access-Control-Allow-Origin")

// 	fmt.Println("Access-Control-Allow-Origin :", val)

// 	correctVal := "*"

// 	if val != correctVal {
// 		t.Fatal("Access-Control-Allow-Origin failed test")
// 	}

// 	val, _ = context.Get("Access-Control-Allow-Methods")

// 	fmt.Println("Access-Control-Allow-Methods: ", val)

// 	correctVal = "POST, GET, OPTIONS, PUT, DELETE"

// 	if val != correctVal {
// 		t.Fatal("Access-Control-Allow-Methods failed test")
// 	}
// }

func initlogging() {

	logfile, err := os.OpenFile("./output_test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
}
