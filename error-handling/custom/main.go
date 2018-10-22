package main

import (
	"fmt"
	"log"
)

func main() {
	i1 := 2.0
	i2 := -2.0

	v, err := Multiply(i2)
	if err != nil {
		//log.Fatalln("Error occurred: ", err) //Prints the error message and exits the program
		log.Println("Error occurred: ", err) //Prints the error message and program continues
		//panic(err) //Prints the error stack and exits the program
	}
	fmt.Println(v)
	fmt.Println("-----")

	v1, err := Multiply(i1)
	if err != nil {
		log.Fatalln("Error occurred: ", err)
	}
	fmt.Println(v1)

}

//Multiply will return the product of the input
func Multiply(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("Pass values greater than 0, passed value:",f)
		return 0, fmt.Errorf("Pass values greater than 0, passed value: %v", f)
	}
	return f * f, nil
}
