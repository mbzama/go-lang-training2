package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")

	if err != nil {
		log.Fatal(err)
	}
	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	//Create slice to hold counts
	buckets := make([]int, 200)

	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	println(word, " - ", bucket)
	return bucket
}
