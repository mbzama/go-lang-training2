package main

import "fmt"

func main() {
	c := make(chan int)

	//Send using named function
	//go foo(c)

	//Send using anonymous function
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	fmt.Println("---")
	//receive
	for v := range c {
		fmt.Println(v)
	}
}

//send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
