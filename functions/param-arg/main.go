package main

import (
	"fmt"
)

func main() {
	greet("Zama", "Md")
	greet("John", "St")
}

func greet(f, l string) {
	fmt.Println(f, l)
}
