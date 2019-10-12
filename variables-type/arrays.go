package main

import "fmt"

func main() {
	var n [100]int /* n is an array of 10 integers */
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 100; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 100; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
