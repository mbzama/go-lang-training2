package main

import (
	"fmt"
)

func main() {
	av1 := average(40, 10, 20, 30, 50)
	fmt.Println("Average:", av1)
	fmt.Println("--------")
	z := []float64{90, 10, 20, 50, 50}
	av2 := average(z...)
	fmt.Println("Average:", av2)
}

func average(sf ...float64) float64 {
	fmt.Println("Inputs:", sf)
	fmt.Printf("Type:%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
