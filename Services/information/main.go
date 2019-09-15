package main

import (
	"log"
	"net/http"
	"projects/Services/common/properties"
	"projects/Services/information/businesslogic"
	"projects/Services/information/sqls"

	"github.com/comail/colog"
)

func main() {
	properties.Init()
	sqls.Init()

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	log.Printf("info: Start Server. listen port --> 8091")

	http.HandleFunc("/getInfoService", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("info: Received request " + "/getInfoService")
		//w.Write([]byte("Called login service."))
		businesslogic.GetInfoService(w, r)
		log.Printf("info: Return to client " + "/getInfoService")

	})

	http.HandleFunc("/editInfoService", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("info: Received request " + "/editInfoService")
		//w.Write([]byte("Called login service."))
		businesslogic.EditInfoService(w, r)
		log.Printf("info: Return to client " + "/editInfoService")

	})
	http.ListenAndServe(":8091", nil)
}
