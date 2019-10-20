package main

import (
	"fmt"
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
)
func main() {

	filename := "config/timonel.yml"
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	yaml, err := simpleyaml.NewYaml(source)
	if err != nil {
		panic(err)
	}

	author, err := yaml.Get("Author").String()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value from yaml file: %s\n", author)

	port, err := yaml.Get("ServerConfig").Get("Port").String()
	//port, err := yaml.Get("ServerConfig").GetIndex(0).String()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value from yaml file: %s\n", port)

	host, err := yaml.Get("ServerConfig").Get("Host").String()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value from yaml file: %s\n", host)

}