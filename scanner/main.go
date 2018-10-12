package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "This is a sequence of words to explore the scanner feature"

	scanner := bufio.NewScanner(strings.NewReader(input))

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Reading input: ", scanner.Err())
	}
}
