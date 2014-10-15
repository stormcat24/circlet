package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func CircletFactory(path string, properties map[string]string) (Circlet, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var replaced = string(data)
	for key, value := range properties {
		replaced = strings.Replace(replaced, "${"+key+"}", value, -1)
	}

	var circlet Circlet
	yaml.Unmarshal([]byte(replaced), &circlet)

	// TODO Validate

	return circlet, nil
}

type Circlet struct {
	Jobs    map[string]CircletJob `yaml:"jobs"`
	Setting CircletSetting        `yaml:"setting"`
}

type CircletOperation interface {
	Execute(jobName string) error
}

func (self *Circlet) Execute(jobName string) error {

	if target, ok := self.Jobs[jobName]; ok {
		fmt.Println("----------[JOB]----------")
		fmt.Printf(" JOB_NAME: %s\n", jobName)
		fmt.Printf(" DESCRIPTION: %s\n", target.Description)
		fmt.Println("")

		api := CircleCIApiFactory(self.Setting, target)
		resp, err := api.ExecuteRequest()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("----------[RESPONSE]----------")
		fmt.Printf(" StatusCode: %d\n", resp.StatusCode)
		for name, value := range resp.Header {
			fmt.Printf(" %s: %s\n", name, strings.Join(value, ","))
		}
		fmt.Println("")

		if resp.StatusCode >= 400 {
			fmt.Println(string(body))
		} else {
			// json
			var jsonBuf bytes.Buffer
			json.Indent(&jsonBuf, body, "", "  ")
			fmt.Println(jsonBuf.String())
		}

		return err
	} else {
		return errors.New(fmt.Sprintf("Not found specified job '%s'.", jobName))
	}
}

type CircletJob struct {
	Description     string            `yaml:"description"`
	Endpoint        string            `yaml:"endpoint"`
	Method          string            `yaml:"method"`
	QueryParameters map[string]string `yaml:"query_parameters"`
	BuildParameters map[string]string `yaml:"build_parameters"`
}

type CircletSetting struct {
	ApiHost  string `yaml:"api_host"`
	ApiToken string `yaml:"api_token"`
}
