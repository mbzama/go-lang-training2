package main

import (
	"fmt"
)

const dollarRate float64 = 71.5

func main() {
	var dollars float64
	fmt.Print("Enter the dollars to convert: ")
	fmt.Scan(&dollars)
	rupees := dollars * dollarRate
	fmt.Println("$", dollars, "with convertion rate(", dollarRate, ") is Rs.", rupees)

}
