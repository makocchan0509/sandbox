package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projects/DogOrCat/web/common/rabbitMQ"
)

func StartWebServer(port string) error {

	log.Println("Running web server ...")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/question", apiMakerHandler(questionHandler))
	http.HandleFunc("/answer", apiMakerHandler(answerHandler))
	http.HandleFunc("/result", apiMakerHandler(resultHandler))
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

type APIResponse struct {
	Result     string `json:"result"`
	ErrMessage string `json:"errMessage"`
}

func answerHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("answer", r.FormValue("answer"))

	var response APIResponse

	err := rabbitMQ.PublishMessage(r.FormValue("answer"))

	if err != nil {
		response = APIResponse{
			Result:     "NG",
			ErrMessage: err.Error(),
		}
	}
	response = APIResponse{
		Result: "OK",
	}

	log.Println(response)
	output, err := json.Marshal(response)

	log.Println(string(output))
	if err != nil {
		log.Println("failed json marchal", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("called resulthandler")
}
