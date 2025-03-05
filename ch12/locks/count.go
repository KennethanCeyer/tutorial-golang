package main

import (
	"fmt"
	"sync"
)


var (
	count int
	mu    sync.Mutex
)


func increment() {
	mu.Lock()
	count++
	mu.Unlock()
}


func main() {
	var wg sync.WaitGroup


	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}


	wg.Wait()
	fmt.Println("Final count:", count)
}
