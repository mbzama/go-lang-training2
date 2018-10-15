package main

//channels BLOCK!!!!
import "fmt"

func main() {
	c := make(chan int)    //receive/send (bidirectional)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Println("----")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
