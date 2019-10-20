package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Author string `yaml:"Author"`
	AppName string `yaml:"AppName"`
	ServerConfig ServerConfig `yaml:"ServerConfig"`
}

type ServerConfig struct {

	Host string `yaml:"Host"`
	Port    string `yaml:"Port"`
}


func main() {
	filename := "config/timonel.yml"
	var config Configuration
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Author)
	fmt.Printf("Value: %#v\n", config.AppName)
	fmt.Printf("Value: %#v\n", config.ServerConfig.Host)
	fmt.Printf("Value: %#v\n", config.ServerConfig.Port)
}