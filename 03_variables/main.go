package main

//Constants
import (
	"fmt"
)

const p string = "COMPLETED"
const (
	A = iota //0
	B        //1
	C        //2
)
const (
	D = iota //0
	E        //1
	F        //2
)

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3*10)
)

func main() {
	fmt.Println(p)
	fmt.Println("---")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("---")
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println("---")
	fmt.Println("binary\t\t\tdecimal")
	fmt.Printf("%b\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t\n", GB)
}
