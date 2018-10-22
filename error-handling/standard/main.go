package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("doesnotexists.txt")

	if err != nil {
		//fmt.Println("Using fmt.Println: ", err)
		//log.Println("Using log.Println: ", err)
		//log.Fatalln("Using log.Fatalln: ", err)
		//log.Panicln("Using log.Fatalln: ", err)
		panic(err)
	}
}
