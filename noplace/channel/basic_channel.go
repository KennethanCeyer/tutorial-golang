package main

import "fmt"

func main() {
	// 최대 2개의 버퍼를 저장할 수 있는
	// 버퍼 채널(Buffered channel)을 생성
	channel := make(chan int, 2)

	// 채널에 값을 보낸다
	channel <- 1
	channel <- 2

	// 채널을 닫음
	close(channel)

	// 가장 먼저 넣은 값부터 채널에서 가져옴 (FIFO)
	fmt.Printf("첫번째 채널: %d\n", <-channel)
	fmt.Printf("두번째 채널: %d\n", <-channel)

	// Zero value 할당
	fmt.Printf("세번째 채널: %d\n", <-channel)
}
