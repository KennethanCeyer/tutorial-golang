package main

import "fmt"

func sendChannel(value int, channel chan<- int) {
	channel <- value
}

func receiveChannel(channel <-chan int) {
	fmt.Printf("채널 수신: %d\n", <-channel)
}

func main() {
	// 양방향(Two-way) 버퍼 채널 할당
	channel := make(chan int, 1)

	sendChannel(1, channel)
	receiveChannel(channel)
}
