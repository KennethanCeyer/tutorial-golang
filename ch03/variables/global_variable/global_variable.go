package main

import "fmt"

var (
	// 전역변수 message
	message = "hello world"
)

func display() {
	fmt.Println(message)
}

func main() {
	// display 함수 호출
	display()
}
