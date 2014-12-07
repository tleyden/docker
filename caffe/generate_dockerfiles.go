package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Params struct {
	GPU            bool
	CPU            bool
	DEVELOP_BRANCH bool
}

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Usage: ./generate_dockerfiles cpu|gpu master|develop")
		return
	}

	params := Params{}

	switch os.Args[1] {
	case "cpu":
		params.CPU = true
		params.GPU = false
	case "gpu":
		params.CPU = false
		params.GPU = true
	default:
		log.Fatal("invalid arg: %v", os.Args[1])
	}

	switch os.Args[2] {
	case "develop":
		params.DEVELOP_BRANCH = true
	case "master":
		params.DEVELOP_BRANCH = false
	default:
		log.Fatal("invalid arg: %v", os.Args[1])

	}

	templateFile := "Dockerfile.template"

	templateBytes, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("docker").Parse(string(templateBytes))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, params)
	if err != nil {
		panic(err)
	}

}
