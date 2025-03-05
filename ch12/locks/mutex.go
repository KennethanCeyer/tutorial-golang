package main

import (
	"fmt"
	"sync"
)


var (
    mu sync.Mutex
    count int
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
    fmt.Println("Count:", count)
}
