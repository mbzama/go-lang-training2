package main

/*
To find out if program has race condition run this command:
	go run -race main.go
	Output: Found 1 data race(s)

Fix the race condition using Atomic
*/
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Go Routines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	fmt.Println("Final Counter:", counter)
}
