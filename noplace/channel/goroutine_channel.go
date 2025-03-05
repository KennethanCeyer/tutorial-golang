package main

import "fmt"

func sendValue(value int, channel chan int) {
	channel <- value
}

func main() {
	// 언버퍼드(Unbuffered) 채널 할당
	unbufferedChannel := make(chan int)

	// 고루틴을 이용해 채널에 1 값을 보냄
	go sendValue(1, unbufferedChannel)

	// 채널에서 값을 가져와 출력
	fmt.Println(<-unbufferedChannel)
}
