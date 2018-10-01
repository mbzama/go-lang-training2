package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("zama", "md"))
	fmt.Println(greet2("zama", "md"))
}

//Return single value
func greet(f, l string) string {
	return fmt.Sprint(f, " ", l)
}

//Return multiple values
func greet2(f, l string) (string, string) {
	return fmt.Sprint(f, " ", l), fmt.Sprint(l, " ", f)
}
