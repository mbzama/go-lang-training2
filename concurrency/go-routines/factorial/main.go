package main

import "fmt"

func main() {
	f := factorial(5)

	for n := range f {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)

	total := 1

	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}
