package main

import "fmt"

func main() {
	//Setup pipeline
	c := gen(1, 2, 3)
	out := sq(c)

	//Consume the output
	for s := range out {
		fmt.Println(s)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
