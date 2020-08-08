package controller

import (
	"html/template"
	"log"
	"net/http"
)

func StartWebServer(port string) error {

	log.Println("Running web server ...")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/question", apiMakerHandler(questionHandler))
	http.HandleFunc("/answer", apiMakerHandler(answerHandler))
	return http.ListenAndServe(":"+port, nil)
}

func apiMakerHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

var templates = template.Must(template.ParseFiles("view/question.html"))

func questionHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "question.html", nil)
	if err != nil {
		log.Println("parse error !!", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("answer", r.FormValue("answer"))

}
