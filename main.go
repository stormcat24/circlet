package main

import (
	"fmt"
	optarg "github.com/jteeuwen/go-pkg-optarg"
	. "github.com/str1ngs/ansi/color"
	os "os"
	"regexp"
	"strings"
)

func main() {

	optarg.Add("c", "config", "circlet configuration file", "circlet.yml")
	optarg.Add("j", "job", "execute job name", nil)
	optarg.Add("p", "props", "write properties", "")

	var configPath, job, properties string

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "c":
			configPath = strings.TrimSpace(opt.String())
		case "j":
			job = strings.TrimSpace(opt.String())
		case "p":
			properties = strings.TrimSpace(opt.String())
		}
	}

	if len(configPath) == 0 {
		fmt.Fprintln(os.Stderr, Red("[ERROR] not specified '-c' option"))
		os.Exit(1)
	}

	if len(job) == 0 {
		fmt.Fprintln(os.Stderr, Red("[ERROR] not specified '-j' option"))
		os.Exit(1)
	}

	tokens := strings.Split(properties, "|")
	matcher := regexp.MustCompile("^(.+)\\s*=\\s*(.+)$")
	propertyMap := make(map[string]string)
	for _, token := range tokens {
		s := strings.TrimSpace(token)
		group := matcher.FindSubmatch([]byte(s))
		if len(group) > 1 {
			if len(group) > 2 {
				propertyMap[string(group[1])] = string(group[2])
			} else {
				propertyMap[string(group[1])] = ""
			}
		}
	}

	// TODO handling validation.
	circlet, _ := CircletFactory(configPath, propertyMap)
	jobError := circlet.Execute(job)
	if jobError != nil {
		fmt.Fprintln(os.Stderr, jobError.Error())
		os.Exit(1)
	} else {
		fmt.Println("Success Job Execution!")
		os.Exit(0)
	}

}
