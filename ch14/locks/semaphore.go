package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)


const maxConcurrency = 3


func worker(id int, sem *semaphore.Weighted) {
	defer sem.Release(1)
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d is done\n", id)
}


func main() {
	sem := semaphore.NewWeighted(maxConcurrency)


	for i := 1; i <= 10; i++ {
		if err := sem.Acquire(context.Background(), 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v\n", err)
			return
		}
		go worker(i, sem)
	}


	// 모든 고루틴이 종료될 때까지 대기
	time.Sleep(10 * time.Second)
}
