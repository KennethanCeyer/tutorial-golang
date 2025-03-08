package main

import (
	"fmt"
	"sync"
	"time"
)


func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d finished\n", id)
}


func main() {
	var wg sync.WaitGroup


	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}


	wg.Wait()
	fmt.Println("All tasks completed")
}
