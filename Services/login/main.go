package main

import (
	"log"
	"net/http"
	"projects/Services/common/properties"
	"projects/Services/login/businesslogic"

	_ "net/http/pprof"

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

	log.Printf("info: Start Server. listen port --> 8090")

	http.HandleFunc("/loginService", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("info: Received request " + "/loginService")
		//w.Write([]byte("Called login service."))
		businesslogic.LoginService(w, r)
		log.Printf("info: Return to client " + "/loginService")
	})

	http.HandleFunc("/getSession", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("info: Received request " + "/getSession")
		businesslogic.GetSession(w, r)
		log.Printf("info: Return to client " + "/getSession")

	})

	http.HandleFunc("/checkSession", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("info: Received request " + "/checkSession")
		businesslogic.CheckSession(w, r)
		log.Printf("info: Return to client " + "/checkSession")

	})

	http.ListenAndServe(":8090", nil)
}
