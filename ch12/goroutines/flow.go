package main

import (
	"fmt"
	"time"
)

// 긴 작업을 수행하는 고루틴
func longRunningTask() {
	for {
		// 무한 루프를 실행
		fmt.Println("Long-running task is still running...")
		time.Sleep(1 * time.Second)
	}
}


func startTask() {
	go longRunningTask() // 긴 작업을 수행하는 고루틴 시작
}


func main() {
	startTask() // 다른 함수에서 긴 작업을 수행하는 고루틴 시작


	// 메인 함수는 2초 후에 종료됨
	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends.")
}
