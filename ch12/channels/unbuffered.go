package main

import "fmt"


func main() {
    // 언버퍼 채널 생성
    unbufferedChannel := make(chan int)


    // 채널에 값을 송신
    unbufferedChannel <- 1


    // 채널에서 값을 수신
    fmt.Println(<-unbufferedChannel)
}
