package main

import "fmt"

// 송신 전용 단방향 채널
func sendChannel(value int, channel chan<- int) {
    channel <- value
}


// 수신 전용 단방향 채널
func receiveChannel(channel <-chan int) {
    fmt.Printf("채널 수신: %d\n", <-channel)
}


func main() {
    // 양방향 채널 생성
    channel := make(chan int, 1)


    // 값 송신
    sendChannel(1, channel)
    
    // 값 수신
    receiveChannel(channel)
}
