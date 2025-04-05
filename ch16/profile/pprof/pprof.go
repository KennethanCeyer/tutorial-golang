package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	_ "net/http/pprof"
)

func denseTask() {
	fmt.Printf("Start denseTask(): %v\n", time.Now())
	data := []int{}
	for i := 0; i < 100000000; i++ {
		data = append(data, rand.Intn(math.MaxInt32))
	}
	fmt.Printf("End denseTask(): %v\n", time.Now())
}

func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			fmt.Printf("pprof server failed: %v\n", err)
		}
	}()

	go denseTask()

	fmt.Println("pprof server started on :6060. Press Ctrl+C to exit.")

	// main 고루틴이 종료되지 않도록 무한 대기 (서버 유지 목적)
	select {}
}
