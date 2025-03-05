package main

import "fmt"


func main() {
    // 최대 2개의 값을 저장할 수 있는 버퍼 채널 생성
    channel := make(chan int, 2)


    // 채널에 값을 송신
    channel <- 1
    channel <- 2


    // 채널을 닫음
    close(channel)


    // 채널에서 값을 수신
    fmt.Printf("첫번째 채널: %d\n", <-channel)
    fmt.Printf("두번째 채널: %d\n", <-channel)


    // 채널이 비어있어도 값을 꺼내면 기본값(Zero value)이 반환됨
    fmt.Printf("세번째 채널: %d\n", <-channel)
}
