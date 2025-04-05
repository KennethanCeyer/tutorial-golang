package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu   sync.Mutex
	data int
)

// writer 고루틴: 주기적으로 data 값을 증가시키며, stop 신호를 받으면 종료
func writer(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // 함수 종료 시 WaitGroup 카운터 감소
	for {
		select {
		case <-stop: // stop 채널이 닫히면 즉시 수신됨 (종료 신호)
			fmt.Println("Writer: Stopping...")
			return // 고루틴 종료
		default: // stop 신호가 없으면 기본 작업 수행
			mu.Lock()
			data++
			fmt.Println("Writer: Incremented data to", data)
			mu.Unlock()
			time.Sleep(150 * time.Millisecond) // 작업 간격 시뮬레이션
		}
	}
}

// reader 고루틴: 주기적으로 data 값을 읽으며, stop 신호를 받으면 종료
func reader(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // 함수 종료 시 WaitGroup 카운터 감소
	for {
		select {
		case <-stop: // stop 채널이 닫히면 즉시 수신됨 (종료 신호)
			fmt.Println("Reader: Stopping...")
			return // 고루틴 종료
		default: // stop 신호가 없으면 기본 작업 수행
			mu.Lock()
			fmt.Println("Reader: Read data", data)
			mu.Unlock()
			time.Sleep(200 * time.Millisecond) // 작업 간격 시뮬레이션
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{}) // 종료 신호용 채널 생성

	// writer와 reader 고루틴 시작
	wg.Add(2) // 2개의 고루틴을 기다리도록 설정
	go writer(stop, &wg)
	go reader(stop, &wg)

	// 고루틴들이 잠시 동안 실행되도록 대기
	time.Sleep(1 * time.Second) // 예: 1초간 실행

	// 모든 고루틴에 종료 신호 전송
	fmt.Println("Main: Sending stop signal...")
	close(stop) // 채널을 닫는 것으로 모든 리스너에게 신호 전달

	// 모든 고루틴(writer, reader)이 실제로 종료될 때까지 대기
	fmt.Println("Main: Waiting for goroutines to stop...")
	wg.Wait()

	fmt.Println("Main: All tasks completed.")
}
