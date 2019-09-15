package businesslogic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Services/common/data"
	"projects/Services/common/net"
	"strings"
)

func GetSession(w http.ResponseWriter, r *http.Request) {

	//Get request parameter
	if r.Method != "POST" {
		log.Println("error: Request is not POST method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("error: Request is not JSON format.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var getSessionReq data.GetSessionReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("ioutil.ReadAll")
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(req, &getSessionReq); err != nil {
		log.Println("json.marshal")
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", getSessionReq)

	sessionId := getSessionReq.SessionId

	//Get session from KVS.
	svo, err := net.GetSession(sessionId)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if svo == "" {
		log.Println("info: Invalid session")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	spValues := strings.Split(svo, ",")

	var uvo data.UserValueObject

	uvo.UserId = spValues[0]
	uvo.UserType = spValues[1]

	//Set content-type on response
	w.Header().Set("Content-Type", "application/json")

	//Create response data
	jr, err := json.Marshal(uvo)

	log.Println("info: response data.>--", string(jr))

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
	return
}

func CheckSession(w http.ResponseWriter, r *http.Request) {

	//Get request parameter
	if r.Method != "POST" {
		log.Println("error: Request is not POST method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("error: Request is not JSON format.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var checkSessionReq data.CheckSessionReq

	//Parse JSON
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("ioutil.ReadAll")
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(req, &checkSessionReq); err != nil {
		log.Println("json.marshal")
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Parsed request JSON parameter
	log.Println("info: Received parameter >--", checkSessionReq)

	sessionId := checkSessionReq.SessionId

	var checkSessionRes data.CheckSessionRes

	//Get session from KVS.
	count, err := net.CheckSession(sessionId)

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if count != 1 {
		log.Println("info: Invalid session")
		w.WriteHeader(http.StatusBadRequest)
		checkSessionRes.Result = "90"
	} else {
		checkSessionRes.Result = "00"
	}

	//Set content-type on response
	w.Header().Set("Content-Type", "application/json")

	//Create response data
	jr, err := json.Marshal(checkSessionRes)

	log.Println("info: response data.>--", string(jr))

	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
	return
}
