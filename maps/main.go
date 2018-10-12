package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}

	for key, val := range myMap {
		fmt.Println(key, " - ", val)
	}
}
