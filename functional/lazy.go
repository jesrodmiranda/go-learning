package main

import "fmt"

func main() {
	fmt.Println(addOrMultiply(true, add(4), multiply(4)))  // 8
	fmt.Println(addOrMultiply(false, add(4), multiply(4))) // 16
}

func add(x int) int {
	fmt.Println("executing add") // this is printed since the functions are evaluated first
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply") // this is printed since the functions are evaluated first
	return x * x
}

func addOrMultiply(add bool, onAdd, onMultiply int) int {
	if add {
		return onAdd
	}
	return onMultiply
}
