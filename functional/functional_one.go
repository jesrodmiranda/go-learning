package main

import "fmt"

func imperative() int {
	n, numElements, s := 1, 0, 0
	for numElements < 10 {
		if n * n % 5 == 0 {
			s += n
			numElements++
		}
		n++
	}
	fmt.Println(s) //275
	return s
}

func main() {

	sum := func (memo interface{}, el interface{}) interface{} {
		return memo.(float64) + el.(float64)
	}
	pred := func (i interface{}) bool {
		return (i.(uint64) * i.(uint64)) % 5 == 0
	}
	createValues := func () []int {
		values := make([]int, 100)
		for num := 1; num <= 100; num++ {
			values = append(values, num)
		}
		return values
	}
	//Reduce(Filter(pred, createValues), sum, int64).(uint64)

	fmt.Printf("Imperative method: %d\n", imperative())
	fmt.Println(sum(float64(5), float64(10)))
	fmt.Println(pred(uint64(5)))
	fmt.Println(createValues())
}