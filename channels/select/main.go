package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)
}

//send
func send(e, o, q chan<- int) {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From eve channel:", v)

		case v := <-o:
			fmt.Println("From odd channel:", v)

		case v := <-q:
			fmt.Println("From quit channel:", v)
			return
		}
	}
}
