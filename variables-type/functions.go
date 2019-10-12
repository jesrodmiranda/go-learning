package main

import "fmt"

func main() {
	var a int = 500
	var b int = 400

	var maxNumber int

	maxNumber = max(a, b)
	fmt.Printf("Maximum Number %d\n", maxNumber)

	x, y := swap("Jesus", "Rodriguez")

	fmt.Printf("Name is %s %s", x , y)
}

func swap (a string, b string) (string, string) {
	return a, b
}

func max(x int, y int) int {
	/* local variable declaration */
	var result int

	if x > y {
		result = x
	} else {
		result = y
	}

	return result
}
