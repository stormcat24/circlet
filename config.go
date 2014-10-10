package main

import (
	"gopkg.in/yaml.v2"
	ioutil "io/ioutil"
)

func ParseCircletYaml(path string) CircletSchema {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var circlet CircletSchema
	yaml.Unmarshal(data, &circlet)

	return circlet
}

type CircletSchema struct {
	Jobs    map[string]CircletJob `yaml:"jobs"`
	Setting CircletSetting        `yaml:"setting"`
}

type CircletJob struct {
	Project     string            `yaml:"project"`
	Description string            `yaml:"description"`
	Parameters  map[string]string `yaml:"parameters"`
}

type CircletSetting struct {
	ApiVersion string `yaml:"api_version"`
	ApiToken   string `yaml:"api_token"`
}
