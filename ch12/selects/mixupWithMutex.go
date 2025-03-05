package main

import (
	"fmt"
	"sync"
	"time"
)

var (
    mu   sync.Mutex
    data int
)

// writer 고루틴: 데이터를 증가시키고 출력한다.
func writer(ch chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ch:
            // 종료 신호를 수신하면 루프를 종료
            return
        default:
            // 데이터를 안전하게 증가시키기 위해 Mutex를 잠금
            mu.Lock()
            data++
            fmt.Println("Data incremented to:", data)
            mu.Unlock()
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// reader 고루틴: 데이터를 읽고 출력한다.
func reader(ch chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ch:
            // 종료 신호를 수신하면 루프를 종료
            return
        default:
            // 데이터를 안전하게 읽기 위해 Mutex를 잠금
            mu.Lock()
            fmt.Println("Data read:", data)
            mu.Unlock()
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan struct{})

    // writer와 reader 고루틴 시작
    wg.Add(2)
    go writer(ch, &wg)
    go reader(ch, &wg)

    // 2초 후에 종료 신호를 채널에 보냄
    time.Sleep(2 * time.Second)
    close(ch)

    // 모든 고루틴이 종료될 때까지 대기
    wg.Wait()
    fmt.Println("All tasks completed")
}
