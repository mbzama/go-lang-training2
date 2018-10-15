package main

//channels BLOCK!!!!
import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
