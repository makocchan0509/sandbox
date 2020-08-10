package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projects/DogOrCat/web/common/dao"
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

type ResultResponse struct {
	Result     string `json:"result"`
	ErrMessage string `json:"errMessage"`
	Total      int    `json:"total"`
	DogPrct    int    `json:"dogprct"`
	CatPrct    int    `json:"catprct"`
}

func resultHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := dao.GetConnectionMysql()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	sql := "SELECT VOTE FROM VOTEBOX"

	rows, err := conn.Query(sql)
	if err != nil {
		log.Printf("Failed query %s", err.Error())
	}

	var vote string
	dogNum := 0
	catNum := 0

	for rows.Next() {
		err := rows.Scan(&vote)
		log.Printf("vote: %s \n", vote)
		if err != nil {
			log.Printf("Failed parse vote %s", err.Error())
		}
		if vote == "dog" {
			dogNum++
		} else if vote == "cat" {
			catNum++
		}
	}

	var res ResultResponse
	var dogPer float64
	var catPer float64

	total := dogNum + catNum
	dogPer = float64(dogNum) / float64(total)
	catPer = float64(catNum) / float64(total)

	res.DogPrct = int(dogPer * 100)
	res.CatPrct = int(catPer * 100)
	res.Total = total
	res.Result = "OK"

	output, err := json.Marshal(res)

	log.Println(string(output))
	if err != nil {
		log.Println("failed json marchal", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
