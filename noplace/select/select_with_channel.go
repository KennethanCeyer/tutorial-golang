package main

import (
	"fmt"
	"sync"
)

func spawn(index int, channel chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		channel <- index*10 + i
	}
	wg.Done()
}

func main() {
	channel := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go spawn(i, channel, wg)
	}
	wg.Wait()

	for i := 0; i < len(channel); i++ {
		fmt.Print("index: %d, value: %d\n", i, channel[i])
	}
}
