package main

import (
	"fmt"
)

// static type declaration
func mainTwo() {
	var a int64
	a = 100
	fmt.Println("x")
	fmt.Printf("a variable is a Type %T", a)
}

func mainOne() {
	var b float64
	c := "50.0"
	fmt.Printf("b variable is a Type %T\n", b)
	fmt.Printf("c variable is a Type %T\n", c)
}

func mainThree() {
	var a, b, c = "Hello", 20.0, 3
	d := 30
	d = 40.0
	e := d
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("a variable is a Type %T\n", a)
	fmt.Printf("b variable is a Type %T\n", b)
	fmt.Printf("c variable is a Type %T\n", c)
}

func mainFour() {
	var a, b string
	var c int
	a = "hello"
	b = "world"
	c = 212 /* Legal */
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func mainFive() {
	// multiply two constants
	const LENGTH int = 20
	const WIDTH int = 30
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("Area is %d\n", area)
}

func main() {
	mainOne()
	mainTwo()
	mainThree()
	mainFour()
	mainFive()
	var a int
	numbers := [6]int{1, 2, 4, 3, 5, 6}

	fmt.Printf("numbers is %T\n", numbers)

	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	/* for loop execution */
	for a = 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	for z := 0; z < 10; z++ {
		fmt.Printf("value of z is %d\n", z)
	}
}
