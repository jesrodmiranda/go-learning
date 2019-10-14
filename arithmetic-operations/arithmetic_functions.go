package main

import (
	"fmt"
)

type Arithmetic interface {
	additionOperation() int
	subtractionOperation() int
	multiplyOperation() int
	divisionOperation() float64
}

type operation struct {
	addend1, addend2 int
}

func (n operation) additionOperation() int {
	return n.addend1 + n.addend2
}

func (n operation) subtractionOperation() int {
	return n.addend1 - n.addend2
}

func (n operation) multiplyOperation() int {
	return n.addend1 * n.addend2
}

func (n operation) divisionOperation() float64 {
	return float64(n.addend1 / n.addend2)
}

func main() {
	var calculate Arithmetic
	fmt.Printf("type of s is %T\n", calculate)
	calculate = operation{7, 2}


	fmt.Printf("type of calculate is %T\n", calculate)
	fmt.Printf("addition: %d\n", calculate.additionOperation())
	fmt.Printf("subtraction: %d\n", calculate.subtractionOperation())
	fmt.Printf("mutiply: %d\n", calculate.multiplyOperation())
	fmt.Printf("division: %.2f\n", calculate.divisionOperation())
}