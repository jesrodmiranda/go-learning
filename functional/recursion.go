package main

import "fmt"

func factorialTailRec(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}
	return factorial(accumulator*val, val-1)
}

func main() {
	fmt.Println(factorialTailRec(12)) // 2432902008176640000
}
