package main

import (
	"log"
	"net/http"
	"projects/Services/common/properties"
	"projects/Services/information/businesslogic"

	"github.com/comail/colog"
)

func main() {
	properties.Init()

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
	http.ListenAndServe(":8091", nil)
}
