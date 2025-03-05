package main

import (
	"fmt"
	"sync"
	"time"
)


func longRunningTask(stop chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-stop:
            // 종료 신호를 받으면 루프를 종료
            fmt.Println("Long-running task is stopping...")
            return
        default:
            // 무한 루프를 실행
            fmt.Println("Long-running task is still running...")
            time.Sleep(1 * time.Second)
        }
    }
}


func startTask(stop chan struct{}, wg *sync.WaitGroup) {
    wg.Add(1)
    go longRunningTask(stop, wg) // 긴 작업을 수행하는 고루틴 시작
}


func main() {
    var wg sync.WaitGroup
    stop := make(chan struct{})


    startTask(stop, &wg) // 다른 함수에서 긴 작업을 수행하는 고루틴 시작


    // 메인 함수는 2초 후에 종료됨
    time.Sleep(2 * time.Second)
    fmt.Println("Main function ends.")


    // 종료 신호를 모든 고루틴에 전송
    close(stop)
    
    // 모든 고루틴이 종료될 때까지 대기
    wg.Wait()
    fmt.Println("All tasks completed")
}
