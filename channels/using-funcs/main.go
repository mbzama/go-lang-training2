package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive using Comma, OK idiom
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

//send
func foo(c chan<- int) {
	c <- 42
}
