package net

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
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
	//Parse response
	r, err = ioutil.ReadAll(resp.Body)

	return r, err
}
