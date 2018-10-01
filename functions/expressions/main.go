package main

import "fmt"

func main() {
	//anonymous function
	greeting := func() {
		fmt.Println("This is anonymous function...")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
