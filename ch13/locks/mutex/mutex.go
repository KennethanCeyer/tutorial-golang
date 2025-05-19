package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func increment() {
	mu.Lock()
	defer mu.Unlock() // Lock 직후 Unlock을 defer로 호출
	count++
	// 만약 여기서 panic이 발생해도 15행에서 defer된 Unlock은 실행됨
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
