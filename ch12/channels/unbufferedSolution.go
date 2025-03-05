package main

import "fmt"


func sendValue(value int, channel chan int) {
    channel <- value
}


func main() {
    // 언버퍼 채널 생성
    unbufferedChannel := make(chan int)


    // 고루틴을 이용해 채널에 값을 송신
    go sendValue(1, unbufferedChannel)


    // 채널에서 값을 수신
    fmt.Println(<-unbufferedChannel)
}
