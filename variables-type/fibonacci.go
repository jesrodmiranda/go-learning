package main

import "fmt"

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func fibonacciIter(n int64) int64 {
	var current int64
	var prev int64
	current, prev = 0, 1
	for i := 0; i < int(n); i++ {
		fmt.Printf("Current: %d Previous: %d\n", current, prev)
		current, prev = current + prev, current
	}
	return current
}

func main() {
	var i int
	for i = 0; i < 50; i++ {
		fmt.Printf("%d ", fibonacciIter(int64(i)))
	}
}