package main

import "fmt"

func main() {
	defer f2() //run before main() exits
	f1()
	f3()
}

func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func f3() {
	fmt.Println("f3")
}
