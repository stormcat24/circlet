package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func CircleCIApiFactory(setting CircletSetting, job CircletJob) CircleCIApi {
	url := fmt.Sprintf("https://%s/api/v1%s", setting.ApiHost, job.Endpoint)

	var queryParams map[string]string
	if job.Parameters != nil {
		queryParams = job.Parameters
	} else {
		queryParams = make(map[string]string)
	}
	queryParams["circle-token"] = setting.ApiToken

	var formData map[string]string
	if job.FormData != nil {
		formData = job.FormData
	} else {
		formData = make(map[string]string)
	}

	return CircleCIApi{url: url, method: job.Method, queryParams: queryParams, formData: formData}
}

type CircleCIApi struct {
	url         string
	method      string
	queryParams map[string]string
	formData    map[string]string
}

type CircleCIApiClient interface {
	ExecuteRequest() (string, error)
}

func (self *CircleCIApi) ExecuteRequest() (*http.Response, error) {

	client := &http.Client{}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("?")

	for key, value := range self.queryParams {
		queryBuffer.WriteString(key)
		queryBuffer.WriteString("=")
		if len(value) > 0 {
			queryBuffer.WriteString(url.QueryEscape(value))
		}
		queryBuffer.WriteString("&")
	}

	requestUrl := self.url + queryBuffer.String()
	fmt.Println("----------[REQUEST]----------")
	fmt.Printf(" URL: %s \n", requestUrl)
	fmt.Printf(" METHOD: %s \n", self.method)
	fmt.Println("")

	req, _ := http.NewRequest(self.method, requestUrl, nil)
	req.Header.Add("User-Agent", "Circlet")
	req.Header.Add("Accept", "application/json")

	// TODO FORM
	return client.Do(req)
}
