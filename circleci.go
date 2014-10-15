package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func CircleCIApiFactory(setting CircletSetting, job CircletJob) CircleCIApi {
	url := fmt.Sprintf("https://%s/api/v1%s", setting.ApiHost, job.Endpoint)

	var queryParams map[string]string
	if job.QueryParameters != nil {
		queryParams = job.QueryParameters
	} else {
		queryParams = make(map[string]string)
	}
	queryParams["circle-token"] = setting.ApiToken

	var buildParams map[string]string
	if job.BuildParameters != nil {
		buildParams = job.BuildParameters
	} else {
		buildParams = make(map[string]string)
	}

	return CircleCIApi{url: url, method: job.Method, queryParams: queryParams, buildParams: buildParams}
}

type CircleCIApi struct {
	url         string
	method      string
	queryParams map[string]string
	buildParams map[string]string
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

	serialized, _ := json.Marshal(self.buildParams)
	json := fmt.Sprintf("{\"build_parameters\": %s}", string(serialized))

	fmt.Println("----------[REQUEST]----------")
	fmt.Printf(" URL: %s \n", requestUrl)
	fmt.Printf(" METHOD: %s \n", self.method)
	fmt.Printf(" DATA: %s \n", json)
	fmt.Println("")

	req, _ := http.NewRequest(self.method, requestUrl, bytes.NewReader([]byte(json)))
	req.Header.Add("User-Agent", "Circlet")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}
