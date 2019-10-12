package main

import "fmt"

func main() {
	type Person struct {
		firstName string
		lastName  string
		fullName  string
		age       int
	}
	var getFullName = func(in Person) string {
		in.fullName = in.firstName + " " + in.lastName
		return in.fullName
	}

	john := Person{
		"john", "doe", "", 30,
	}

	fmt.Println(getFullName(john))
	fmt.Println(john)
}
