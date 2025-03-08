package main

import (
    "fmt"
    "net/http"
    "time"
    _ "net/http/pprof"
)

func optimizedCommunication() {
    ch := make(chan int, 100000) // 버퍼가 있는 채널 사용

    for i := 0; i < 100000; i++ {
        go func(n int) {
            time.Sleep(1 * time.Millisecond) // 인위적인 지연
            ch <- n
        }(i)
    }

    for i := 0; i < 100000; i++ {
        fmt.Println(<-ch) // 수신된 값을 출력
    }
}

func main() {
    go optimizedCommunication()
    http.ListenAndServe("localhost:6060", nil)
}

