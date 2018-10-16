package main

/*
Example source: https://austburn.me/blog/a-better-fan-in-fan-out-example.html
*/
import (
	"fmt"
	"time"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)

	itemToShip := []Item{
		Item{0, "Shirt", 1 * time.Second},
		Item{1, "TV", 5 * time.Second},
		Item{2, "Computer", 5 * time.Second},
		Item{3, "Phone", 2 * time.Second},
	}

	go func() {
		for _, item := range itemToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()

	return items
}

func PackItems(done <-chan bool, items <-chan Item) <-chan int {
	packages := make(chan int)

	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Shipping package no: %d\n", item.ID)
			}
		}
		close(packages)
	}()
	return packages
}

func main() {
	done := make(chan bool)
	defer close(done)

	start := time.Now()

	packages := PackItems(done, PrepareItems(done))

	numPackages := 0

	for range packages {
		numPackages++
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
