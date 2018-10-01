package main

import "fmt"

func main() {
	a := 43
	addr := &a

	fmt.Println("a : ", a)
	fmt.Println("Address of a : ", addr)
	fmt.Println("Value at address Address of a : ", *addr)
	fmt.Println("----")

	var b *int = &a
	fmt.Println("Address of b : ", b)
	fmt.Println("Value at address Address of b : ", *b)

	*b = 42

	fmt.Println("a : ", a)
}
