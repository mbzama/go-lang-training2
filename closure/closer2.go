package main

import "fmt"

var x int

func main() {
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("------")

	incw := wrapper()
	fmt.Println(incw())
	fmt.Println(incw())
}

func increment() int {
	x++
	return x
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
