package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Logging(path string, logFileName string) {

	logFullPath := path + "/" + logFileName

	logfile, err := os.OpenFile(logFullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("fail open logfile")
		os.Exit(1)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
