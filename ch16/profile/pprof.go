package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"
)

func denseTask(wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Start denseTask(): %v\n", time.Now())

    data := []int{}
    for i := 0; i < 100000000; i++ {
        data = append(data, rand.Intn(math.MaxInt32))
        time.Sleep(time.Nanosecond)
    }

    fmt.Printf("End denseTask(): %v\n", time.Now())
}

func main() {
    var wg sync.WaitGroup

    go func() {
        fmt.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // wg에 2개의 Task를 등록하여 끝나지 않도록 한다
    // denseTask의 작업은 일정 개수 이상 도달하면 종료되지만,
    // main() 고루틴은 종료되지 않으므로 CPU와 메모리 상태를 조회할 수 있다.
    wg.Add(1)
    wg.Add(1)

    go denseTask(&wg)
    wg.Wait()
}
