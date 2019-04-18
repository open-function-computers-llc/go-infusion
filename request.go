package infusion

import (
	"bytes"
	"net/http"
)

var infusionBaseURL = "https://api.infusionsoft.com/crm/rest/v1"
var c http.Client
var authToken string
var refreshToken string

func getRequest(route string) (*http.Response, error) {
	url := infusionBaseURL + route
	log.Info("GET: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	return client.Do(req)
}

func postRequest(route string, body []byte) (*http.Response, error) {
	url := infusionBaseURL + route
	log.Info("POST: "+url, "POSTBODY: "+string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
