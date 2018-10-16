package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 5, 9)

	//FAN OUT
	//Distribute the square work to 2 go routines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	//FAN IN
	//Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
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

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("Type: %T \n", cs)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
