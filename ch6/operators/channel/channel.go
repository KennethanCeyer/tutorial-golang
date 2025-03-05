package main

import "fmt"

func main() {
    // 채널 선언
    messages := make(chan string)

    // 고루틴을 통해 메시지 전송
    go func() {
        messages <- "hello"
    }()

    // 채널에서 메시지 수신
    msg := <-messages
    fmt.Println(msg) // hello
}
