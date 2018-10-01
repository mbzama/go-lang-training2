package main

import "fmt"

func main() {
	a := 43
	addr := &a

	fmt.Println("a : ", a)
	fmt.Println("Address of a : ", addr)
	fmt.Println("Value at address Address of a : ", *addr)

}
