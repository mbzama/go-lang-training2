package main

import (
	"fmt"

	"github.com/mbzama/go-lang-training2/02_package/stringutil"
	"github.com/segmentio/fasthash/fnv1a"
)

func main() {
	fmt.Println(stringutil.MyName + "\n")
	fmt.Printf(stringutil.Reverse("oG olleH") + "\n")

	h1 := fnv1a.HashString64("Hello World!")
	fmt.Println("FNV-1a hash of 'Hello World!':", h1)
}
