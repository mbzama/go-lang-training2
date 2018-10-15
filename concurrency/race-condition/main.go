package main

/*
To find out if program has race condition run this command:
	go run -race main.go
	Output: Found 1 data race(s)
*/
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
