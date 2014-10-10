package main

import (
	"fmt"
	optarg "github.com/jteeuwen/go-pkg-optarg"
	. "github.com/str1ngs/ansi/color"
	os "os"
	"strings"
)

func init() {
	fmt.Printf("### %s ###\n", Green("Circlet is support tool of CircleCI."))
}

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

	// TODO parse properties string
	// TODO execute specified job
	fmt.Println(properties)

	result := ParseCircletYaml(configPath)
	fmt.Println(result)

}
