package net

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"projects/Services/common/data"
	"projects/Services/common/properties"
	"strings"
)

func JsonPostRequestSender(url string, input []byte) (r []byte, err error) {

	//Create Http request
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}
	// Set Content-Type
	req.Header.Set("Content-Type", "application/json")

	//POST request to login service.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	//Parse JSON  response
	r, err = ioutil.ReadAll(resp.Body)

	return r, err
}

func QueryPostRequestSender(url string, requestValue url.Values) (r []byte, err error) {

	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(requestValue.Encode()),
	)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	//Parse JSON response
	r, err = ioutil.ReadAll(resp.Body)

	return r, err
}

func GetSessionByJson(sessionId string) (r []byte, err error) {

	prop := properties.GetProp()
	url := prop.Session.Url

	log.Println("info: URL>--", url)

	reqData := new(data.GetSessionReq)

	reqData.SessionId = sessionId

	log.Println("info: reqData>--", reqData)

	input, err := json.Marshal(reqData)

	log.Println("info: input>--", input)

	//Create Http request
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}
	// Set Content-Type
	req.Header.Set("Content-Type", "application/json")

	//POST request to login service.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	//Parse JSON  response
	r, err = ioutil.ReadAll(resp.Body)

	return r, err

}
