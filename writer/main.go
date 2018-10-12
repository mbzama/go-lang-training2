package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Printing to standard output...")
	fmt.Fprintln(os.Stdout, "Printing to standard output...")
	io.WriteString(os.Stdout, "Printing to standard output...")
}
