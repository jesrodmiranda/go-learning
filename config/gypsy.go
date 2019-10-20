package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func nodeToMap(node yaml.Node) yaml.Map {
	m, ok := node.(yaml.Map)
	if !ok {
		panic(fmt.Sprintf("%v is not of type map", node))
	}
	return m
}

func nodeToList(node yaml.Node) (yaml.List) {
	m, ok := node.(yaml.List)
	if !ok {
		panic(fmt.Sprintf("%v is not of type list", node))
	}
	return m
}

func main() {
	// filename := os.Args[1]
	filename := "config/config.yml"
	file, err := yaml.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	value := nodeToList(nodeToMap(file.Root)["bar"])[0]
	fmt.Printf("Value: %#v\n", value)
}